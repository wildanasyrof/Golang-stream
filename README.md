# Golang Fiber API with Uber Fx 🚀

This is a scalable **Golang REST API** built with [Fiber](https://gofiber.io/) and [Uber Fx](https://fx.uber.org/) for dependency injection. It follows **Clean Architecture** principles and supports **JWT authentication**, **MySQL database**, and **modular components**.

## 📌 Features

- 🔹 User Authentication (Register, Login, JWT)
- 🔹 Role-Based Access Control (RBAC)
- 🔹 Anime & Category Management (Many-to-Many)
- 🔹 Fiber Web Framework for High Performance
- 🔹 Uber Fx for Dependency Injection
- 🔹 GORM ORM with MySQL Integration
- 🔹 Standardized API Responses & Request Validation
- 🔹 Docker Support for Easy Deployment

## 📂 Project Structure

```
GOLANG-STREAMING-API/
│── api/
│   ├── handler/       # Business logic for request handling
│   ├── middleware/    # Middleware functions (e.g., authentication, logging)
│   ├── routes/        # API route definitions
│
│── cmd/               # Application entry point
│
│── config/            # Configuration files
│
│── internal/          # Internal application logic
│   ├── dto/           # Data Transfer Objects
│   ├── entity/        # Database entity definitions
│   ├── module/        # Module-level components
│   ├── repository/    # Database queries and interactions
│   ├── service/       # Business logic services
│
│── migration/         # Database migration files
│
│── pkg/               # Utility packages
│
│── .env               # Environment variables
│── .gitignore         # Git ignored files
│── go.mod             # Go module dependencies
│── go.sum             # Dependency checksums
```

## 🚀 Getting Started

### 1️⃣ Prerequisites
- Install **Go** (>=1.18)
- Install **MySQL**

### 2️⃣ Setup & Installation

```sh
# Clone the repository
git clone https://github.com/yourusername/golang-fiber-app.git
cd golang-fiber-app

# Install dependencies
go mod tidy
```

### 3️⃣ Configure Environment
Create a `.env` file in the root directory:

```
APP_PORT=3000
DB_USER=root
DB_PASSWORD=yourpassword
DB_NAME=golang_fiber_db
DB_HOST=localhost
DB_PORT=3306
JWT_SECRET=your_secret_key
```

### 4️⃣ Run Migrations
```sh
go run cmd/migrate.go up
```

### 5️⃣ Start the Server
```sh
go run cmd/main.go
```

Server runs at **`http://localhost:3000`** 🚀

## 🛠 API Endpoints

### 🔐 Auth Routes
| Method | Endpoint         | Description       |
|--------|-----------------|-------------------|
| POST   | `/auth/register` | Register a user  |
| POST   | `/auth/login`    | Login & get JWT  |

### 👤 User Routes (Auth Required)
| Method | Endpoint   | Description       |
|--------|-----------|-------------------|
| GET    | `/user`   | Get user profile  |
| PUT    | `/user`   | Update profile    |

### 🎬 Anime Routes (Admin Only)
| Method | Endpoint      | Description         |
|--------|--------------|---------------------|
| POST   | `/anime`     | Create new anime    |
| GET    | `/anime`     | Get all anime       |

### 🎭 Category Routes
| Method | Endpoint         | Description          |
|--------|-----------------|----------------------|
| GET    | `/category`      | Get all categories  |
| POST   | `/category` (Admin) | Create category |


## 📌 Contributing
Feel free to fork and submit pull requests! 🚀

---

Made with ❤️ using **Golang**, **Fiber**, and **Uber Fx**.

