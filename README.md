# Todo App - Go Clean Architecture

A production-ready Go web application built with clean architecture principles, featuring user authentication and todo management APIs.

## Features

- ✅ **Clean Architecture**: Domain, Application, Infrastructure layers
- ✅ **RESTful API**: Versioned endpoints (`/api/v1/`)
- ✅ **JWT Authentication**: Secure user authentication
- ✅ **Todo Management**: Create, Read, Update, Delete operations
- ✅ **MongoDB Integration**: Document database with proper indexing
- ✅ **Input Validation**: Comprehensive request validation
- ✅ **Error Handling**: Standardized error responses with error codes
- ✅ **Health Checks**: `/health` endpoint for monitoring
- ✅ **Logging**: Structured logging with correlation IDs
- ✅ **Unit & Integration Tests**: Comprehensive test suite (80%+ coverage)
- ✅ **Docker Support**: Docker and Docker Compose configurations
- ✅ **API Documentation**: Swagger/OpenAPI support

## Project Structure

```
todo-app/
├── cmd/
│   ├── migrate/          # Database migration runner
│   ├── server/           # Main application entry point
│   └── worker/           # Background job worker
├── internal/
│   ├── features/         # Feature modules (auth, todo)
│   │   ├── auth/         # Authentication feature
│   │   └── todo/         # Todo feature
│   ├── infrastructure/   # External integrations (DB, cache)
│   ├── shared/           # Shared utilities (middleware, config, response)
│   └── services/         # Business services
├── pkg/                  # Public packages (logger, errors, utils)
├── tests/
│   ├── unit/             # Unit tests
│   ├── integration/      # Integration tests
│   └── helpers/          # Test helpers and mocks
├── deployments/          # Docker and Kubernetes configs
├── docs/                 # API documentation
└── scripts/              # Utility scripts
```

## Getting Started

### Prerequisites

- Go 1.25.6 or higher
- MongoDB 4.4+
- Docker & Docker Compose (optional)
- Make (optional, for using Makefile commands)

### Installation

1. **Clone the repository**
   ```bash
   git clone <repo-url>
   cd todo-app
   ```

2. **Setup environment variables**
   ```bash
   cp .env.example .env
   # Edit .env with your configuration
   ```

3. **Install dependencies**
   ```bash
   make deps
   ```

### Running the Application

#### Option 1: Local Development
```bash
# Build and run
make run

# Or run directly
go run ./cmd/server/main.go
```

#### Option 2: Docker
```bash
# Start MongoDB and application
make docker-up

# View logs
make docker-logs

# Stop containers
make docker-down
```

## API Endpoints

### Base URL
```
http://localhost:8080/api/v1
```

### Health Check
```
GET /health
```

### Authentication

#### Sign Up
```
POST /auth/signup
Content-Type: application/json

{
  "email": "user@example.com",
  "password": "SecurePassword123",
  "name": "John Doe"
}

Response: 201 Created
{
  "data": {
    "token": "eyJhbGc...",
    "email": "user@example.com",
    "name": "John Doe"
  }
}
```

#### Login
```
POST /auth/login
Content-Type: application/json

{
  "email": "user@example.com",
  "password": "SecurePassword123"
}

Response: 200 OK
{
  "data": {
    "token": "eyJhbGc...",
    "email": "user@example.com",
    "name": "John Doe"
  }
}
```

### Todo Operations

All todo endpoints require JWT authentication header:
```
Authorization: Bearer <token>
```

#### Create Todo
```
POST /todos
Content-Type: application/json

{
  "title": "Learn Go",
  "description": "Master Go programming language",
  "completed": false
}

Response: 201 Created
{
  "data": {
    "id": "507f1f77bcf86cd799439011",
    "title": "Learn Go",
    "description": "Master Go programming language",
    "completed": false,
    "created_at": "2024-01-20T12:00:00Z",
    "updated_at": "2024-01-20T12:00:00Z"
  }
}
```

#### Get All Todos
```
GET /todos

Response: 200 OK
{
  "data": [
    {
      "id": "507f1f77bcf86cd799439011",
      "title": "Learn Go",
      ...
    }
  ]
}
```

#### Get Todo by ID
```
GET /todos/:id

Response: 200 OK
{
  "data": {
    "id": "507f1f77bcf86cd799439011",
    ...
  }
}
```

#### Update Todo
```
PUT /todos/:id
Content-Type: application/json

{
  "title": "Learn Go Advanced",
  "description": "Master advanced Go topics",
  "completed": true
}

Response: 200 OK
{
  "data": {
    "id": "507f1f77bcf86cd799439011",
    ...
  }
}
```

#### Delete Todo
```
DELETE /todos/:id

Response: 204 No Content
```

## Configuration

### Environment Variables
```
# Server
PORT=8080

# Database
MONGO_URI=mongodb://localhost:27017
DATABASE_NAME=todos
COLLECTION_NAME=todos

# JWT
JWT_SECRET=your-secret-key-here

# Environment
NODE_ENV=development
```

## Testing

### Run All Tests
```bash
make test
```

### Run Unit Tests Only
```bash
make test-unit
```

### Run Integration Tests Only
```bash
make test-integration
```

### Generate Coverage Report
```bash
make test-coverage
# Opens coverage.html in browser
```

### Check Coverage Percentage
```bash
make test-coverage-check
```

## Code Quality

### Format Code
```bash
make fmt
```

### Run Linter
```bash
make lint
```

### Run Go Vet
```bash
make vet
```

## Building

### Build Binary
```bash
make build
```

### Build Docker Image
```bash
make docker-build
```

## Error Handling

The API returns standardized error responses:

```json
{
  "code": "VALIDATION_ERROR",
  "message": "Invalid email format",
  "details": "Email must be a valid email address"
}
```

### Error Codes
- `VALIDATION_ERROR` (400) - Input validation failed
- `UNAUTHORIZED` (401) - Authentication required or failed
- `FORBIDDEN` (403) - Access denied
- `NOT_FOUND` (404) - Resource not found
- `CONFLICT` (409) - Resource already exists
- `INTERNAL_ERROR` (500) - Server error

## Development

### Directory Structure Best Practices

1. **Features**: Keep domain logic in `internal/features/`
2. **Shared Code**: Common utilities in `internal/shared/`
3. **Infrastructure**: Database, cache implementations in `internal/infrastructure/`
4. **Tests**: Parallel structure under `tests/` directory
5. **Packages**: Reusable packages in `pkg/`

### Adding a New Feature

1. Create feature directory: `internal/features/feature-name/`
2. Structure:
   ```
   feature-name/
   ├── dto/           # Data transfer objects
   ├── entity/        # Domain entities
   ├── handler/       # HTTP handlers
   ├── repository/    # Data access
   ├── routes/        # Route definitions
   └── usecase/       # Business logic
   ```

3. Register in `internal/shared/config/app.go`
4. Add tests in `tests/unit/` and `tests/integration/`

## Deployment

### Docker Compose
```bash
docker-compose up -d
```

### Kubernetes
```bash
kubectl apply -f deployments/kubernetes/
```

### Environment-Specific Configuration
- Development: `.env`
- Production: Set environment variables before running

## Performance Considerations

- **Connection Pooling**: MongoDB connection pooling configured
- **Indexing**: Database indexes on frequently queried fields
- **Caching**: Ready for Redis integration
- **Rate Limiting**: Middleware available for rate limiting

## Security Features

- ✅ JWT token validation
- ✅ Input validation and sanitization
- ✅ CORS configuration
- ✅ Request logging and audit trails
- ✅ Error message sanitization (no internal details exposed)

## Contributing

1. Create feature branch: `git checkout -b feature/name`
2. Make changes and test: `make test`
3. Run code quality checks: `make fmt lint`
4. Commit with descriptive message
5. Push and create pull request

## License

MIT License - see LICENSE file for details

## Support

For issues and questions:
1. Check existing GitHub issues
2. Create new issue with detailed description
3. Include error logs and reproduction steps

---

**Last Updated**: January 20, 2026
**Status**: Production Ready
