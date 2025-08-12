# 🎬 Full-Stack Movie Application

A **complete full-stack microservices application** with authentication, JWT-based security, and persistent data storage — all containerized with Docker Compose.

Built using **Node.js**, **Go**, **PostgreSQL**, and a **vanilla HTML/JS frontend**.

---

## 🏗️ Architecture Overview

```mermaid
graph TD
    Browser[💻 Browser] -- HTTP Requests --> NodeBackend[🟢 Node.js User Service]
    NodeBackend -- Serves --> Frontend[🌐 HTML/JS Frontend]
    Frontend -- Register/Login --> NodeBackend
    Frontend -- Authenticated Requests (JWT) --> GoBackend[💛 Go Movie Service]
    NodeBackend -- DB Connection --> PostgreSQL[🗄 PostgreSQL Database]
    GoBackend -- DB Connection --> PostgreSQL
```

**Services:**

* **Frontend:** Lightweight HTML + JS (no frameworks) served by Node.js.
* **User Service (Node.js):** Handles registration/login, password hashing, JWT issuance.
* **Movie Service (Go):** Requires valid JWT, manages favorite movies.
* **PostgreSQL:** Shared database for both services.

---

## 🚀 Tech Stack

| Layer               | Technologies                                            |
| ------------------- | ------------------------------------------------------- |
| **Frontend**        | HTML5, CSS3, JavaScript (ES6+), Fetch API               |
| **Backend (User)**  | Node.js, Express.js, bcryptjs, jsonwebtoken, pg, dotenv |
| **Backend (Movie)** | Go, Gorilla Mux, JWT-Go, pq, cors                       |
| **Database**        | PostgreSQL                                              |
| **DevOps**          | Docker, Docker Compose                                  |

---

## ⚙️ Setup & Run

### 1️⃣ Clone the repository

```bash
git clone <repository_url>
cd <repository_directory>
```

### 2️⃣ Build & start with Docker Compose

```bash
docker-compose up --build
```

This will:

* Build Node.js & Go services
* Start PostgreSQL
* Run services on:

  * **Node.js:** `http://localhost:3100`
  * **Go:** `http://localhost:4100`

### 3️⃣ Open in browser

```
http://localhost:3100
```

---

## 🎯 Features

✅ **User Registration & Login** with password hashing (bcrypt)
✅ **JWT Authentication** for secure API requests
✅ **Add & View Favorite Movies** tied to the logged-in user
✅ **Persistent Storage** in PostgreSQL
✅ **Containerized Microservices** with Docker Compose

---

## 🖥 Usage Flow

1. **Register** → Create an account with email & password.
2. **Login** → Get a JWT token stored in `localStorage`.
3. **Add Movies** → Enter a movie name and add it to your list.
4. **View Movies** → See all your saved favorites instantly.
5. **Logout** → Clear session & token.

---

## 📂 Project Structure

```
📦 project-root
 ┣ 📂 frontend       # Static HTML/JS files
 ┣ 📂 user-service   # Node.js authentication service
 ┣ 📂 movie-service  # Go movie management service
 ┣ 📂 db             # Database initialization scripts
 ┣ 📜 docker-compose.yml
 ┗ 📜 README.md
```

---

## 🐳 Docker Compose Services

| Service           | Port | Description                 |
| ----------------- | ---- | --------------------------- |
| **user-service**  | 3100 | Node.js User API + frontend |
| **movie-service** | 4100 | Go Movie API                |
| **db**            | 5432 | PostgreSQL database         |

---

## 📸 Demo Preview


<img width="811" height="751" alt="1" src="https://github.com/user-attachments/assets/a664b83e-122a-4e1b-b6a0-35f57d94e5d3" />
<img width="811" height="751" alt="2" src="https://github.com/user-attachments/assets/27ac6745-6e5c-4594-a62f-64c7da736d7f" />
<img width="811" height="814" alt="3" src="https://github.com/user-attachments/assets/40b891d8-3256-4f56-b25a-b5750295dea4" />
<img width="811" height="814" alt="4" src="https://github.com/user-attachments/assets/ec740573-c841-48c7-aaba-e8da6691c0f9" />
<img width="809" height="543" alt="5" src="https://github.com/user-attachments/assets/b951bdd8-28e0-4ec9-859b-5e775b242ddd" />
<img width="809" height="543" alt="6" src="https://github.com/user-attachments/assets/c16652dc-ef22-43a5-a67d-3a1466068f82" />
<img width="808" height="606" alt="7" src="https://github.com/user-attachments/assets/2b6e3945-87be-49e5-8a9c-0e1a46090391" />
<img width="808" height="606" alt="8" src="https://github.com/user-attachments/assets/b26b040e-e217-4297-b1f3-f59ec47d972f" />
<img width="808" height="659" alt="9" src="https://github.com/user-attachments/assets/9cc9833d-25e6-47f7-aa17-c9754e1e85d8" />


---



💡 **Pro Tip:** This project is a great starting point for learning **microservices, JWT authentication, and container orchestration**.

---

