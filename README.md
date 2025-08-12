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

*(Optional: Add screenshots or a GIF here to showcase the app in action)*

<img width="808" height="659" alt="Screenshot from 2025-08-12 16-42-56" src="https://github.com/user-attachments/assets/1c034492-7652-473e-a0f2-70b00e93076c" />
<img width="808" height="606" alt="Screenshot from 2025-08-12 16-42-42" src="https://github.com/user-attachments/assets/dce819c9-db53-48d7-afca-bb661391e245" />
<img width="808" height="606" alt="Screenshot from 2025-08-12 16-41-55" src="https://github.com/user-attachments/assets/17f93db3-fa8f-49b2-ba73-a7c1a11ae8a8" />
<img width="809" height="543" alt="Screenshot from 2025-08-12 16-41-43" src="https://github.com/user-attachments/assets/76a1d34f-8806-4a55-8ed3-2bf683970674" />
<img width="809" height="543" alt="Screenshot from 2025-08-12 16-41-33" src="https://github.com/user-attachments/assets/b27e52c3-c18c-4918-803b-00978fef34bf" />
<img width="811" height="814" alt="Screenshot from 2025-08-12 16-41-17" src="https://github.com/user-attachments/assets/ef072727-84db-4e8a-8ee4-bded256aecd3" />
<img width="811" height="814" alt="Screenshot from 2025-08-12 16-40-59" src="https://github.com/user-attachments/assets/1f80f58d-39ec-4de3-83bb-4e9dd4c7f323" />
<img width="811" height="751" alt="Screenshot from 2025-08-12 16-40-45" src="https://github.com/user-attachments/assets/d908f5b9-d2c2-4b5e-97b4-c4dea5ea9cae" />
<img width="811" height="751" alt="Screenshot from 2025-08-12 16-39-30" src="https://github.com/user-attachments/assets/f4fbf7b9-9e05-4391-ad8c-c1b43de672c7" />



---

💡 **Pro Tip:** This project is a great starting point for learning **microservices, JWT authentication, and container orchestration**.

---

