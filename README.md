# Go Fiber JWT Example

This repository contains an example project demonstrating how to use JWT (JSON Web Tokens) with the Go Fiber web framework.

## Project Structure

```
go-fiber-jwt-example/
├── config/
│   ├── config.go
├── controller/
│   ├── auth.go
│   ├── book.go
├── database/
│   ├── db.go
├── middleware/
│   ├── auth.go
├── model/
│   ├── book.go
│   ├── user.go
├── router/
│   ├── router.go
├── utils/
│   ├── password.go
│   ├── token.go
└── main.go
```


## Getting Started

### Prerequisites

- Go 1.16 or higher
- A configured Go environment

### Installation

1. Clone the repository:

    ```sh
    git clone https://github.com/yourusername/go-fiber-jwt-example.git
    cd go-fiber-jwt-example
    ```

2. Install dependencies:

    ```sh
    go mod tidy
    ```

### Running the Application

1. Start the application:

    ```sh
    go run main.go
    ```

2. The server will start on `http://localhost:3000`.

### Project Structure

- [`config/`](config): Contains the application configuration.
- [`controller/`](controller)
  - [`auth.go`](controller/auth.go): Contains the authentication handlers.
  - [`book.go`](controller/book.go): Contains the book handlers.
- [`database/`](database): Contains the database connection logic.
- [`middleware/`](middleware)
  - [`jwt.go`](middleware/jwt.go): Ensures authenticated access to routes by validating and decoding JSON Web Tokens (JWT).
- [`model/`](model): Contains the data models.
- [`router/`](router): Contains the application routes.
- [`utils/`](utils)
  - [`password_generator.go`](utils/password_generator.go): Contains the `GeneratePassword` and `ComparePassword` function.
  - [`token_generator`](utils/token_generator.go): Contains the `GenerateToken` and `VerifyToken` function.
- [`main.go`](main.go): The entry point of the application.

## License

This project is licensed under the MIT License. See the LICENSE file for details.