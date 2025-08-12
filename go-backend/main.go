package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/rs/cors"
)

type UserClaims struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
	jwt.StandardClaims
}

type Movie struct {
	ID       int    `json:"id"`
	UserID   int    `json:"user_id"`
	MovieName string `json:"movie_name"`
}

var db *sql.DB
var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

func init() {
	// Load .env file if not in production
	if os.Getenv("APP_ENV") != "production" {
		err := godotenv.Load()
		if err != nil {
			log.Println("Error loading .env file, assuming environment variables are set.")
		}
	}

	jwtSecret = []byte(os.Getenv("JWT_SECRET"))
	if len(jwtSecret) == 0 {
		log.Fatal("JWT_SECRET environment variable not set.")
	}

	var err error
	connStr := os.Getenv("DATABASE_URL")
	if len(connStr) == 0 {
		log.Fatal("DATABASE_URL environment variable not set.")
	}

	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error opening database connection: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	log.Println("Successfully connected to PostgreSQL!")
}

func main() {
	router := mux.NewRouter()

	// CORS setup
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3100"}, // Allow Node.js frontend
		AllowCredentials: true,
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Authorization", "Content-Type"},
	})

	router.Use(c.Handler)

	// Handle OPTIONS requests for CORS preflight
	router.Methods("OPTIONS").HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	// Protected routes
	router.HandleFunc("/movies", authMiddleware(addMovie)).Methods("POST")
	router.HandleFunc("/movies", authMiddleware(getMovies)).Methods("GET")

	port := os.Getenv("PORT")
	if port == "" {
		port = "4100" // Default port
	}

	log.Printf("Go backend listening on port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

func authMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header required", http.StatusUnauthorized)
			return
		}

		tokenString := strings.Replace(authHeader, "Bearer ", "", 1)
		token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
			// Validate the alg is what you expect: `HS256`
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return jwtSecret, nil
		})

		if err != nil {
			http.Error(w, "Invalid or expired token: "+err.Error(), http.StatusUnauthorized)
			return
		}

		if claims, ok := token.Claims.(*UserClaims); ok && token.Valid {
			// Add user ID to request context for downstream handlers
			r.Header.Set("X-User-ID", fmt.Sprintf("%d", claims.ID))
			next.ServeHTTP(w, r)
		} else {
			http.Error(w, "Invalid token claims", http.StatusUnauthorized)
		}
	}
}

func addMovie(w http.ResponseWriter, r *http.Request) {
	userIDStr := r.Header.Get("X-User-ID")
	if userIDStr == "" {
		http.Error(w, "User ID not found in context", http.StatusInternalServerError)
		return
	}

	var movie Movie
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&movie); err != nil {
		log.Printf("Error decoding movie payload: %v", err)
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if movie.MovieName == "" {
		log.Println("Movie name is empty in payload")
		http.Error(w, "Movie name cannot be empty", http.StatusBadRequest)
		return
	}

	var userID int
	_, err := fmt.Sscanf(userIDStr, "%d", &userID)
	if err != nil {
		http.Error(w, "Invalid User ID format", http.StatusInternalServerError)
		return
	}

	movie.UserID = userID

	// Insert the new movie
	_, err = db.Exec("INSERT INTO movies (user_id, movie_name) VALUES ($1, $2)", movie.UserID, movie.MovieName)
	if err != nil {
		log.Printf("Error inserting movie: %v", err)
		http.Error(w, "Failed to add movie", http.StatusInternalServerError)
		return
	}

	// Fetch and return the updated list of movies for the user
	getMoviesForUser(w, userID)
}

func getMovies(w http.ResponseWriter, r *http.Request) {
	userIDStr := r.Header.Get("X-User-ID")
	if userIDStr == "" {
		http.Error(w, "User ID not found in context", http.StatusInternalServerError)
		return
	}

	var userID int
	_, err := fmt.Sscanf(userIDStr, "%d", &userID)
	if err != nil {
		http.Error(w, "Invalid User ID format", http.StatusInternalServerError)
		return
	}

	getMoviesForUser(w, userID)
}

func getMoviesForUser(w http.ResponseWriter, userID int) {
	rows, err := db.Query("SELECT id, user_id, movie_name FROM movies WHERE user_id = $1 ORDER BY id DESC", userID)
	if err != nil {
		log.Printf("Error querying movies for user %d: %v", userID, err)
		http.Error(w, "Failed to retrieve movies", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	movies := []Movie{}
	for rows.Next() {
		var m Movie
		if err := rows.Scan(&m.ID, &m.UserID, &m.MovieName); err != nil {
			log.Printf("Error scanning movie row: %v", err)
			continue
		}
		movies = append(movies, m)
	}

	if err = rows.Err(); err != nil {
		log.Printf("Error iterating movie rows: %v", err)
		http.Error(w, "Failed to retrieve movies", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}
