# Book Management API

This is a book management API built with Golang and Echo framework. It supports CRUD operations for books and categories, authentication using JWT, and role-based access control.

## Features
- User authentication (JWT-based login and registration)
- Role-based access control (Admin, Member)
- CRUD operations for books
- CRUD operations for categories
- CRUD operations for favorites
- CRUD operations for reviews
- Search & Filtering for books
- Middleware for protected routes
- PostgreSQL integration with GORM
- Configuration management using Viper
- Logging with Zap

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
- **Search & Filter Books**: `GET /books?title=Go&category=Programming&rating=4&page=1&limit=10&sort=title`

### Categories
- **Create Category**: `POST /categories`
- **Get All Categories**: `GET /categories`
- **Get Category by ID**: `GET /categories/:id`
- **Update Category**: `PUT /categories/:id`
- **Delete Category**: `DELETE /categories/:id`

### Favorites
- **Add to Favorites**: `POST /favorites`
- **Get User Favorites**: `GET /favorites`
- **Remove from Favorites**: `DELETE /favorites/:id`

### Reviews
- **Create Review**: `POST /reviews`
- **Get Reviews for a Book**: `GET /reviews/:book_id`
- **Update Review**: `PUT /reviews/:id`
- **Delete Review**: `DELETE /reviews/:id`

## Role-Based Access Control
- Users have roles (`admin`, `member`).
- Certain endpoints (e.g., creating/deleting books) are restricted to `admin` users.
- Use JWT token for authentication and role validation.

## Authentication
- The API uses JWT for authentication.
- To access protected routes, include the token in the `Authorization` header as: `Bearer <your-token>`.

## License
MIT License

