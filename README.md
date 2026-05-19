# Go Backend Boilerplate

Production-ready Golang backend boilerplate using Echo, PostgreSQL, JWT Authentication, RBAC, Docker, Swagger, and Clean Architecture style.

---

## Features

- REST API with Echo
- PostgreSQL with GORM
- JWT Authentication
- Register & Login
- Role-Based Access Control (RBAC)
- Admin-only User Management
- Pagination
- Search & Filter Users
- Soft Delete
- Request Validation
- Custom Error Handling
- Auto Migration
- Admin Seeder
- Graceful Shutdown
- Docker Support
- Swagger Documentation

---

## Tech Stack

- Go
- Echo
- PostgreSQL
- GORM
- JWT
- Docker
- Swagger

---

## Default Admin Account

```txt
email: admin@example.com
password: admin123

# Go Backend Boilerplate

Production-ready Golang backend boilerplate using Echo, PostgreSQL, JWT Authentication, RBAC, Docker, Swagger, and Clean Architecture style.

---

## Features

- REST API with Echo
- PostgreSQL with GORM
- JWT Authentication
- Register & Login
- Role-Based Access Control (RBAC)
- Admin-only User Management
- Pagination
- Search & Filter Users
- Soft Delete
- Request Validation
- Custom Error Handling
- Auto Migration
- Admin Seeder
- Graceful Shutdown
- Docker Support
- Swagger Documentation

---

## Tech Stack

- Go
- Echo
- PostgreSQL
- GORM
- JWT
- Docker
- Swagger

---

## Default Admin Account

```txt
email: admin@example.com
password: admin123

Run locally:
cp .env.example .env

go mod tidy

go run cmd/api/main.go

Run with Docker:
docker compose up --build
