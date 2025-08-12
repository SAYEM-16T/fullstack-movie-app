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

<img width="808" height="659" alt="9" src="https://github.com/user-attachments/assets/e3b2e290-eb54-41bf-bc84-ce19b0261291" />
<img width="808" height="606" alt="8" src="https://github.com/user-attachments/assets/0f2edc4b-b480-4c26-b8d2-7c6c986b2423" />
<img width="808" height="606" alt="7" src="https://github.com/user-attachments/assets/747e8168-7516-43d7-8350-121296792812" />
<img width="809" height="543" alt="6" src="https://github.com/user-attachments/assets/257571ec-c802-4dc3-b92b-6b5f8b79cbde" />
<img width="809" height="543" alt="5" src="https://github.com/user-attachments/assets/8277df86-e05e-449e-a4f8-ff5e29728f92" />
<img width="811" height="814" alt="4" src="https://github.com/user-attachments/assets/9085d880-df72-4a22-8880-ea09209ac0a2" />
<img width="811" height="814" alt="3" src="https://github.com/user-attachments/assets/c7467b16-4472-48f4-afdc-0a62bd0f82f8" />
<img width="811" height="751" alt="2" src="https://github.com/user-attachments/assets/5a508b30-edfd-41a8-bfab-47cf756f4648" />
<img width="811" height="751" alt="1" src="https://github.com/user-attachments/assets/9a805a9a-9dc0-43f5-bd38-e8d730cce77f" />




---

💡 **Pro Tip:** This project is a great starting point for learning **microservices, JWT authentication, and container orchestration**.

---

