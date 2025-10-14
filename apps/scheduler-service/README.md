# Job-Scheduler (Golang)

## 🔧 Getting Started

### Prerequisites

- Go 1.23+
- PostgreSQL (or compatible database)

### Running Locally
- Run docker-compose.yaml file, it will create a required postgres database container.
```bash
go mod tidy
go run main.go
```

## 📚 Documentation

API documentation is automatically generated through the Huma framework and available at the `/docs` endpoint when the server is running.

## 🧪 Testing

Run the tests with:

```bash
go test ./...
```