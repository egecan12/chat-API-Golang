# 🐹 Golang Chat Application with Supabase 🐘

![Supabase](https://img.shields.io/badge/Supabase-3ECF8E?style=for-the-badge&logo=supabase&logoColor=white)
![Golang](https://img.shields.io/badge/Go-00ADD8?style=for-the-badge&logo=go&logoColor=white)

A real-time chat application built using **Golang** for the backend and **Supabase** PostgreSQL as the database. This project includes user registration, login with JWT authentication, and WebSocket-based real-time messaging.

---

## 🚀 Features

- **User Authentication**:
  - Registration with hashed passwords.
  - Login with JWT-based token generation.
- **Real-time Messaging**:
  - WebSocket-based message handling.
- **Supabase Integration**:
  - PostgreSQL hosted on Supabase for scalable database solutions.
- **Secure Backend**:
  - Password hashing with bcrypt.
  - Token-based authentication using JWT.

---

## 📋 Prerequisites

Before you start, make sure you have the following installed:

1. **Golang**:
   - Install Go from [https://go.dev/](https://go.dev/).
   - Verify installation:
     ```bash
     go version
     ```
2. **Supabase Account**:
   - Sign up at [https://supabase.com/](https://supabase.com/).
   - Create a new project and set up your PostgreSQL database.
3. **Database Schema**:
   - Create the following tables in Supabase using the SQL Editor:
     ```sql
     CREATE TABLE users (
         id SERIAL PRIMARY KEY,
         username TEXT NOT NULL UNIQUE,
         password TEXT NOT NULL,
         created_at TIMESTAMP DEFAULT now()
     );
     ```

---

## 🛠️ Setup Instructions

Follow these steps to get the application running:

### 1. Clone the Repository

```bash
git clone https://github.com/your-username/go-chat-app.git
cd go-chat-app
```

### 2. Install Dependencies

```bash
go mod tidy
```

### 3. Create a `.env` File

Add your Supabase connection details and JWT secret:

```
DATABASE_URL=postgres://<username>:<password>@<host>:5432/<database>
JWT_SECRET=your_super_secret_key
```

### 4. Run the Application

```bash
go run cmd/api/main.go
```

The server will start on `http://localhost:8080`.

---

## 📚 API Endpoints

### 1. **Register User**

- **Endpoint**: `POST /register`
- **Request Body**:
  ```json
  {
    "username": "example_user",
    "password": "secure_password"
  }
  ```
- **Response**:
  ```json
  {
    "message": "User registered successfully!"
  }
  ```

### 2. **Login User**

- **Endpoint**: `POST /login`
- **Request Body**:
  ```json
  {
    "username": "example_user",
    "password": "secure_password"
  }
  ```
- **Response**:
  ```json
  {
    "token": "your_jwt_token"
  }
  ```

### 3. **Protected Route Example**

- **Endpoint**: `GET /protected`
- **Headers**:
  ```json
  {
    "Authorization": "Bearer your_jwt_token"
  }
  ```
- **Response**:
  ```json
  {
    "message": "Welcome to the protected route!"
  }
  ```

---

## 🌐 Supabase Configuration

### **Steps to Get Your Supabase Connection String**

1. Log in to your Supabase account.
2. Select your project.
3. Navigate to **Settings > Database**.
4. Copy the connection string in the format:
   ```
   postgres://<username>:<password>@<host>:5432/<database>
   ```

---

## 🛠️ Project Structure

```
go-chat-app/
├── cmd/
│   └── api/
│       └── main.go       # Main entry point of the application
├── pkg/
│   ├── db/
│   │   └── db.go         # Database connection logic
│   ├── routes/
│   │   └── routes.go     # API routes definition
│   └── handlers/
│       └── user.go       # User registration and login logic
└── .env                  # Environment variables (not included in repo)
```

---

## 🤝 Contributing

1. Fork the repository.
2. Create a feature branch:
   ```bash
   git checkout -b feature/your-feature
   ```
3. Commit your changes:
   ```bash
   git commit -m "Add your message here"
   ```
4. Push to your branch:
   ```bash
   git push origin feature/your-feature
   ```
5. Open a pull request.

---

## 🧑‍💻 Authors

- **Your Name**: [GitHub Profile](https://github.com/your-username)

---

## 📜 License

This project is licensed under the MIT License. See the `LICENSE` file for details.
