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

- [`config/`](command:_github.copilot.openRelativePath?%5B%7B%22scheme%22%3A%22file%22%2C%22authority%22%3A%22%22%2C%22path%22%3A%22%2FUsers%2Fcakici%2FDesktop%2Fgo-fiber-jwt-example%2Fconfig%2F%22%2C%22query%22%3A%22%22%2C%22fragment%22%3A%22%22%7D%5D "/Users/cakici/Desktop/go-fiber-jwt-example/config/"): Configuration files.
- [`controller/`](command:_github.copilot.openRelativePath?%5B%7B%22scheme%22%3A%22file%22%2C%22authority%22%3A%22%22%2C%22path%22%3A%22%2FUsers%2Fcakici%2FDesktop%2Fgo-fiber-jwt-example%2Fcontroller%2F%22%2C%22query%22%3A%22%22%2C%22fragment%22%3A%22%22%7D%5D "/Users/cakici/Desktop/go-fiber-jwt-example/controller/"): Handlers for different routes.
    - [`auth.go`](command:_github.copilot.openSymbolFromReferences?%5B%22auth.go%22%2C%5B%7B%22uri%22%3A%7B%22%24mid%22%3A1%2C%22fsPath%22%3A%22%2FUsers%2Fcakici%2FDesktop%2Fgo-fiber-jwt-example%2FReadme.md%22%2C%22external%22%3A%22file%3A%2F%2F%2FUsers%2Fcakici%2FDesktop%2Fgo-fiber-jwt-example%2FReadme.md%22%2C%22path%22%3A%22%2FUsers%2Fcakici%2FDesktop%2Fgo-fiber-jwt-example%2FReadme.md%22%2C%22scheme%22%3A%22file%22%7D%2C%22pos%22%3A%7B%22line%22%3A7%2C%22character%22%3A8%7D%7D%5D%5D "Go to definition"): Authentication-related handlers.
    - [`book.go`](command:_github.copilot.openSymbolFromReferences?%5B%22book.go%22%2C%5B%7B%22uri%22%3A%7B%22%24mid%22%3A1%2C%22fsPath%22%3A%22%2FUsers%2Fcakici%2FDesktop%2Fgo-fiber-jwt-example%2FReadme.md%22%2C%22external%22%3A%22file%3A%2F%2F%2FUsers%2Fcakici%2FDesktop%2Fgo-fiber-jwt-example%2FReadme.md%22%2C%22path%22%3A%22%2FUsers%2Fcakici%2FDesktop%2Fgo-fiber-jwt-example%2FReadme.md%22%2C%22scheme%22%3A%22file%22%7D%2C%22pos%22%3A%7B%22line%22%3A8%2C%22character%22%3A8%7D%7D%5D%5D "Go to definition"): Book-related handlers.
- [`database/`](command:_github.copilot.openRelativePath?%5B%7B%22scheme%22%3A%22file%22%2C%22authority%22%3A%22%22%2C%22path%22%3A%22%2FUsers%2Fcakici%2FDesktop%2Fgo-fiber-jwt-example%2Fdatabase%2F%22%2C%22query%22%3A%22%22%2C%22fragment%22%3A%22%22%7D%5D "/Users/cakici/Desktop/go-fiber-jwt-example/database/"): Database connection and setup.
- [`middleware/`](command:_github.copilot.openSymbolFromReferences?%5B%22middleware%2F%22%2C%5B%7B%22uri%22%3A%7B%22%24mid%22%3A1%2C%22fsPath%22%3A%22%2FUsers%2Fcakici%2FDesktop%2Fgo-fiber-jwt-example%2FReadme.md%22%2C%22external%22%3A%22file%3A%2F%2F%2FUsers%2Fcakici%2FDesktop%2Fgo-fiber-jwt-example%2FReadme.md%22%2C%22path%22%3A%22%2FUsers%2Fcakici%2FDesktop%2Fgo-fiber-jwt-example%2FReadme.md%22%2C%22scheme%22%3A%22file%22%7D%2C%22pos%22%3A%7B%22line%22%3A11%2C%22character%22%3A4%7D%7D%5D%5D "Go to definition"): Middleware functions.
    - [`auth.go`](command:_github.copilot.openSymbolFromReferences?%5B%22auth.go%22%2C%5B%7B%22uri%22%3A%7B%22%24mid%22%3A1%2C%22fsPath%22%3A%22%2FUsers%2Fcakici%2FDesktop%2Fgo-fiber-jwt-example%2FReadme.md%22%2C%22external%22%3A%22file%3A%2F%2F%2FUsers%2Fcakici%2FDesktop%2Fgo-fiber-jwt-example%2FReadme.md%22%2C%22path%22%3A%22%2FUsers%2Fcakici%2FDesktop%2Fgo-fiber-jwt-example%2FReadme.md%22%2C%22scheme%22%3A%22file%22%7D%2C%22pos%22%3A%7B%22line%22%3A7%2C%22character%22%3A8%7D%7D%5D%5D "Go to definition"): JWT authentication middleware.
- [`model/`](command:_github.copilot.openRelativePath?%5B%7B%22scheme%22%3A%22file%22%2C%22authority%22%3A%22%22%2C%22path%22%3A%22%2FUsers%2Fcakici%2FDesktop%2Fgo-fiber-jwt-example%2Fmodel%2F%22%2C%22query%22%3A%22%22%2C%22fragment%22%3A%22%22%7D%5D "/Users/cakici/Desktop/go-fiber-jwt-example/model/"): Data models.
    - [`book.go`](command:_github.copilot.openSymbolFromReferences?%5B%22book.go%22%2C%5B%7B%22uri%22%3A%7B%22%24mid%22%3A1%2C%22fsPath%22%3A%22%2FUsers%2Fcakici%2FDesktop%2Fgo-fiber-jwt-example%2FReadme.md%22%2C%22external%22%3A%22file%3A%2F%2F%2FUsers%2Fcakici%2FDesktop%2Fgo-fiber-jwt-example%2FReadme.md%22%2C%22path%22%3A%22%2FUsers%2Fcakici%2FDesktop%2Fgo-fiber-jwt-example%2FReadme.md%22%2C%22scheme%22%3A%22file%22%7D%2C%22pos%22%3A%7B%22line%22%3A8%2C%22character%22%3A8%7D%7D%5D%5D "Go to definition"): Book model.
    - [`user.go`](command:_github.copilot.openSymbolFromReferences?%5B%22user.go%22%2C%5B%7B%22uri%22%3A%7B%22%24mid%22%3A1%2C%22fsPath%22%3A%22%2FUsers%2Fcakici%2FDesktop%2Fgo-fiber-jwt-example%2FReadme.md%22%2C%22external%22%3A%22file%3A%2F%2F%2FUsers%2Fcakici%2FDesktop%2Fgo-fiber-jwt-example%2FReadme.md%22%2C%22path%22%3A%22%2FUsers%2Fcakici%2FDesktop%2Fgo-fiber-jwt-example%2FReadme.md%22%2C%22scheme%22%3A%22file%22%7D%2C%22pos%22%3A%7B%22line%22%3A15%2C%22character%22%3A8%7D%7D%5D%5D "Go to definition"): User model.
- [`router/`](command:_github.copilot.openRelativePath?%5B%7B%22scheme%22%3A%22file%22%2C%22authority%22%3A%22%22%2C%22path%22%3A%22%2FUsers%2Fcakici%2FDesktop%2Fgo-fiber-jwt-example%2Frouter%2F%22%2C%22query%22%3A%22%22%2C%22fragment%22%3A%22%22%7D%5D "/Users/cakici/Desktop/go-fiber-jwt-example/router/"): Route definitions.
- [`utils/`](command:_github.copilot.openSymbolFromReferences?%5B%22utils%2F%22%2C%5B%7B%22uri%22%3A%7B%22%24mid%22%3A1%2C%22fsPath%22%3A%22%2FUsers%2Fcakici%2FDesktop%2Fgo-fiber-jwt-example%2FReadme.md%22%2C%22external%22%3A%22file%3A%2F%2F%2FUsers%2Fcakici%2FDesktop%2Fgo-fiber-jwt-example%2FReadme.md%22%2C%22path%22%3A%22%2FUsers%2Fcakici%2FDesktop%2Fgo-fiber-jwt-example%2FReadme.md%22%2C%22scheme%22%3A%22file%22%7D%2C%22pos%22%3A%7B%22line%22%3A18%2C%22character%22%3A4%7D%7D%5D%5D "Go to definition"): Utility functions.
    - [`password.go`](command:_github.copilot.openSymbolFromReferences?%5B%22password.go%22%2C%5B%7B%22uri%22%3A%7B%22%24mid%22%3A1%2C%22fsPath%22%3A%22%2FUsers%2Fcakici%2FDesktop%2Fgo-fiber-jwt-example%2FReadme.md%22%2C%22external%22%3A%22file%3A%2F%2F%2FUsers%2Fcakici%2FDesktop%2Fgo-fiber-jwt-example%2FReadme.md%22%2C%22path%22%3A%22%2FUsers%2Fcakici%2FDesktop%2Fgo-fiber-jwt-example%2FReadme.md%22%2C%22scheme%22%3A%22file%22%7D%2C%22pos%22%3A%7B%22line%22%3A19%2C%22character%22%3A8%7D%7D%5D%5D "Go to definition"): Password hashing and verification.
    - [`token.go`](command:_github.copilot.openSymbolFromReferences?%5B%22token.go%22%2C%5B%7B%22uri%22%3A%7B%22%24mid%22%3A1%2C%22fsPath%22%3A%22%2FUsers%2Fcakici%2FDesktop%2Fgo-fiber-jwt-example%2FReadme.md%22%2C%22external%22%3A%22file%3A%2F%2F%2FUsers%2Fcakici%2FDesktop%2Fgo-fiber-jwt-example%2FReadme.md%22%2C%22path%22%3A%22%2FUsers%2Fcakici%2FDesktop%2Fgo-fiber-jwt-example%2FReadme.md%22%2C%22scheme%22%3A%22file%22%7D%2C%22pos%22%3A%7B%22line%22%3A20%2C%22character%22%3A8%7D%7D%5D%5D "Go to definition"): JWT token generation and validation.
- [`main.go`](command:_github.copilot.openRelativePath?%5B%7B%22scheme%22%3A%22file%22%2C%22authority%22%3A%22%22%2C%22path%22%3A%22%2FUsers%2Fcakici%2FDesktop%2Fgo-fiber-jwt-example%2Fmain.go%22%2C%22query%22%3A%22%22%2C%22fragment%22%3A%22%22%7D%5D "/Users/cakici/Desktop/go-fiber-jwt-example/main.go"): Entry point of the application.

## Contributing

1. Fork the repository.
2. Create a new branch (`git checkout -b feature-branch`).
3. Make your changes.
4. Commit your changes (`git commit -m 'Add some feature'`).
5. Push to the branch (`git push origin feature-branch`).
6. Create a new Pull Request.

## License

This project is licensed under the MIT License. See the LICENSE file for details.