# Go Gin Template

A clean and modular template for building microservices with Golang, Gin, PostgreSQL, and Docker. This template follows Clean Architecture principles, making it easy to extend, test, and maintain.

## Features

- **Gin**: Fast, minimalist web framework for building REST APIs.
- **PostgreSQL**: Reliable and feature-rich open-source database.
- **Docker**: Easily containerize your application for consistent development and production environments.
- **Clean Architecture**: Separation of concerns for scalable and maintainable code.

## Getting Started

### Prerequisites

- [Go](https://golang.org/dl/) 1.20+
- [Docker](https://www.docker.com/get-started)
- [Docker Compose](https://docs.docker.com/compose/install/)

### Installation

1. Clone the repository:

   ```bash
   git clone https://github.com/yourusername/go-gin-template.git
   cd go-gin-template
   ```
2. Create a `.env` file based on the example `.env.example`:

   ```bash
   cp .env.example .env
   ```
3. Build and run the application with Docker:

   ```bash
   docker-compose up --build
   ```
4. Access the application at `http://localhost:8080`.

### Project Structure

The project is organized based on Clean Architecture principles, separating the business logic from delivery and data access layers.

- **`internal/domain/`**: Contains core business logic and entities.
- **`internal/usecase/`**: Application use cases or services that interact with domain models.
- **`internal/repository/`**: Implements data access and repository patterns.
- **`internal/delivery/`**: Handles incoming requests, includes HTTP handlers, routes, and middleware.
- **`internal/config/`**: Manages configuration and environment variables.
- **`pkg/`**: Contains shared code like logging utilities.

### Usage

- **Endpoints**: Modify or extend handlers in `internal/delivery/http/handler.go`.
- **Database Migrations**: Manage database schema changes in the `InitDB` folder (to be created if needed).

### Environment Variables

- `DB_HOST`: Database host (default: `localhost`)
- `DB_PORT`: Database port (default: `5432`)
- `DB_USER`: Database user (default: `postgres`)
- `DB_PASSWORD`: Database password (default: `password`)
- `DB_NAME`: Database name (default: `template_db`)

### Contributing

Contributions are welcome! Feel free to submit issues and pull requests to improve the template.

### License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

**Happy coding! 🚀**
