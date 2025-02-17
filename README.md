# Book API with Golang and Echo

This is a Book Management API built with Golang using the Echo framework. It follows a clean architecture with separate layers for handlers, services, and repositories. The database is PostgreSQL, and the project uses Viper for configuration management.

## Features

- CRUD operations for books
- Category management
- Retrieve books with their category names
- PostgreSQL integration with GORM
- Configuration management using Viper
- Structured logging with Zap
- Authentication using JWT (Planned)
- Unit tests (Planned)
- Docker support (Planned)
- API documentation with Swagger (Planned)

## Project Structure

```
/book-api
├── cmd
│   └── main.go
├── config
│   └── config.yaml
├── internal
│   ├── handlers
│   ├── models
│   ├── repositories
│   ├── services
│   └── database
├── go.mod
└── README.md
```

## Getting Started

### Prerequisites

- Golang (>=1.18)
- PostgreSQL
- Docker (Optional)

### Installation

1. Clone the repository:
   ```sh
   git clone https://github.com/pinklaosa/book-api.git
   cd book-api
   ```
2. Install dependencies:
   ```sh
   go mod tidy
   ```
3. Set up PostgreSQL using Docker:
   ```sh
   docker-compose up -d
   ```
4. Update `config/config.yaml` with your database credentials.

5. Run the project:
   ```sh
   go run cmd/main.go
   ```

### API Endpoints

#### Books
| Method | Endpoint       | Description             |
|--------|---------------|-------------------------|
| GET    | /books        | Get all books          |
| GET    | /books/:id    | Get book by ID         |
| POST   | /books        | Create a new book      |
| PUT    | /books/:id    | Update book by ID      |
| DELETE | /books/:id    | Delete book by ID      |

#### Categories
| Method | Endpoint        | Description              |
|--------|----------------|--------------------------|
| GET    | /categories    | Get all categories       |
| POST   | /categories    | Create a new category    |

## Environment Configuration

The project uses Viper for configuration. Modify `config/config.yaml`:
```yaml
app:
  port: 8080

database:
  host: localhost
  port: 5432
  user: myuser
  password: mypassword
  dbname: mydatabase
  sslmode: disable
```

## Future Improvements

- Implement JWT authentication
- Add unit tests
- Improve error handling
- Dockerize the application
- Integrate Swagger for API documentation

## License

This project is licensed under the MIT License. Feel free to use and modify it!

---

💡 **Contributions are welcome!** If you find issues or want to add features, feel free to open a pull request.

