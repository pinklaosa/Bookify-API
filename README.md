# Book Management API

This is a book management API built with Golang and Echo framework. It follows a layered architecture consisting of handlers, services, and repositories. It supports CRUD operations for books and categories, and includes authentication using JWT.

## Features
- User authentication (JWT-based login and registration)
- CRUD operations for books
- CRUD operations for categories
- Middleware for protected routes
- PostgreSQL integration with GORM
- Configuration management using Viper

## Project Structure
```
/book-api
├── cmd
│   └── main.go
├── internal
│   ├── handlers      # Handles HTTP requests
│   ├── services      # Business logic layer
│   ├── repositories  # Database interaction layer
│   ├── models        # Database models
│   └── middleware    # Middleware for authentication & logging
└── go.mod
```

## Installation
1. Clone the repository:
   ```sh
   git clone <repository-url>
   cd book-api
   ```
2. Install dependencies:
   ```sh
   go mod tidy
   ```
3. Set up environment variables in `config.yaml`:
   ```yaml
   server:
     port: 8080
   database:
     host: localhost
     port: 5432
     user: myuser
     password: mypassword
     dbname: mydatabase
   jwt:
     secret_key: "your-secret-key"
   ```
4. Start the PostgreSQL database using Docker:
   ```sh
   docker-compose up -d
   ```
5. Run the application:
   ```sh
   air  # If using Air for live reload
   ```

## API Endpoints

### Authentication
- **Register**: `POST /auth/register`
- **Login**: `POST /auth/login`

### Books
- **Create Book**: `POST /books`
- **Get All Books**: `GET /books`
- **Get Book by ID**: `GET /books/:id`
- **Update Book**: `PUT /books/:id`
- **Delete Book**: `DELETE /books/:id`

### Categories
- **Create Category**: `POST /categories`
- **Get All Categories**: `GET /categories`
- **Get Category by ID**: `GET /categories/:id`
- **Update Category**: `PUT /categories/:id`
- **Delete Category**: `DELETE /categories/:id`

## Authentication
- The API uses JWT for authentication.
- To access protected routes, include the token in the `Authorization` header as: `Bearer <your-token>`.


