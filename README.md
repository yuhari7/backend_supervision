# Super Vision Assessment Test

## Description

Proyek ini adalah pembuatan mini aplikasi untuk manage article

## Features

- Clean Architecture (separated layers for controllers, services, repositories, etc.)
- Article management system (CRUD operations)
- JWT authentication for secure endpoints
- Modular project structure for scalability
- Microservice-oriented design

## Technologies Used

- **Backend**: Go (Golang)
- **Web Framework**: Echo
- **Database**: PostgreSQL - User & MySQL - Article
- **Authentication**: JWT
- **ORM**: GORM
- **Frontend**: Next.js (with Tailwind CSS for styling)

## Setup Instructions

### Prerequisites

Make sure you have the following installed:

- Go (Golang) 1.18+
- PostgreSQL and MySQL
- Any other dependencies you might need for your environment

### Install Dependencies

1. Clone the repository:

   ```bash
   git clone https://github.com/yourusername/yourproject.git
   cd yourproject
   ```

2. Install Go dependencies:

```go
go mod tidy
```

3. Setup Database and Migrations

```bash - untuk article
migrate -path ./migrations -database "mysql://root:@tcp(localhost:3306)/article_service" up
```

```bash - untuk user
migrate -database "postgres://user:pass@localhost:5432/your_db?sslmode=disable" -path ./migrations up

```

4. Running Applications

```go - user
go run cmd/main.go
```

```go - article
go run cmd/main.go
```

## Documentation

https://www.postman.com/EhCTkZy2TTNgHwF/workspace/super-vision-api-demo/collection/7847915-9852d153-37dd-49ba-83a4-9bbe1816e7eb?action=share&creator=7847915

## untuk Front End

Running secara lokal

1. Clone Source Code
2. Install depedencies

```bash
npm install
```

3. Running application

```bash
npm run dev
```
