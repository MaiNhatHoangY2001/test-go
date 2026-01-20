# Contributing Guide

Thank you for your interest in contributing to Todo App! This guide will help you get started.

## Code of Conduct

- Be respectful and inclusive
- Provide constructive feedback
- Focus on ideas, not individuals
- Follow the project's coding standards

---

## Getting Started

### 1. Fork and Clone
```bash
# Fork the repository on GitHub
# Clone your fork
git clone https://github.com/your-username/todo-app.git
cd todo-app
```

### 2. Set Up Development Environment
```bash
# Install dependencies
make deps

# Start MongoDB
docker run -d -p 27017:27017 mongo:latest

# Run tests to verify setup
make test
```

### 3. Create a Feature Branch
```bash
git checkout -b feature/your-feature-name
```

---

## Development Workflow

### Code Style

Follow Go best practices:
- Use `gofmt` for formatting
- Use meaningful variable names
- Add comments for exported functions
- Keep functions small and focused

### Running Code Quality Checks

```bash
# Format code
make fmt

# Run linter
make lint

# Run go vet
make vet

# Run all checks before commit
make all
```

### Writing Tests

- Write unit tests for all business logic
- Use `testify/mock` for mocking dependencies
- Aim for 80%+ coverage
- Tests should be independent and fast

```go
// Good test structure
func (suite *TodoHandlerTestSuite) TestCreateTodo_Success() {
    // Arrange - setup test data and mocks
    mockRepo.On("Create", mock.Anything).Return(nil)
    
    // Act - execute the function
    result := handler.CreateTodo(input)
    
    // Assert - verify the result
    assert.NoError(suite.T(), result)
}
```

### Commit Guidelines

```bash
# Use descriptive commit messages
git commit -m "feat: add todo filtering by status

- Implement filter logic in usecase
- Add tests for filter scenarios
- Update API documentation"

# Commit message format:
# <type>: <subject>
# 
# <body>
# 
# <footer>

# Types: feat, fix, docs, style, refactor, test, chore
```

### Create a Pull Request

1. Push to your fork
   ```bash
   git push origin feature/your-feature-name
   ```

2. Create PR with detailed description:
   - **What**: What does this PR do?
   - **Why**: Why is this change needed?
   - **How**: How does it work?
   - **Testing**: How to test it?

3. Link related issues: `Closes #123`

---

## Adding a New Feature

### Step 1: Design the Feature
- Plan the feature structure
- Think about edge cases
- Consider backward compatibility

### Step 2: Create Feature Directory
```
internal/features/feature-name/
├── dto/              # Data transfer objects
├── entity/           # Domain entities
├── handler/          # HTTP handlers
├── repository/       # Data access interfaces
├── routes/           # Route definitions
└── usecase/          # Business logic
```

### Step 3: Implement Feature
```
1. Create entities (models)
2. Create DTOs (request/response)
3. Implement repository interface
4. Implement use cases
5. Create handlers
6. Register routes
```

### Step 4: Write Tests
- Unit tests for use cases
- Handler tests with mocks
- Integration tests
- Aim for >80% coverage

### Step 5: Update Documentation
- Add API documentation
- Update README if needed
- Add configuration docs
- Add example requests

### Step 6: Register in App Config
Update `internal/shared/config/app.go`:

```go
// Initialize repositories
featureCollection := client.Database(...).Collection(...)
featureRepo := infrastructure.NewMongoFeatureRepository(featureCollection)

// Initialize use cases
featureUseCase := usecase.NewFeatureUseCase(featureRepo)

// Initialize handler
featureHandler := handler.NewFeatureHandler(featureUseCase, logger)

// Register routes
v1 := router.Group("/api/v1")
routes.SetupFeatureRoutes(v1, featureHandler)
```

---

## Testing Guidelines

### Unit Tests
```bash
make test-unit
```

### Integration Tests
```bash
make test-integration
```

### Check Coverage
```bash
make test-coverage-check
```

### Coverage Report
```bash
make test-coverage
# Opens coverage.html
```

---

## Repository Rules

### Directory Structure
Keep the clean architecture:
- `features/` - Domain features
- `shared/` - Shared utilities
- `infrastructure/` - External integrations
- `pkg/` - Public packages
- `tests/` - Test files

### Do's ✅
- Write clear, descriptive commit messages
- Create tests for new code
- Update documentation
- Follow Go conventions
- Keep functions focused
- Add error handling

### Don'ts ❌
- Mix concerns in handlers
- Skip error handling
- Add println/log statements
- Hardcode values
- Create duplicate code
- Skip tests

---

## Release Process

1. **Update version** in appropriate files
2. **Update CHANGELOG.md**
3. **Create release PR**
4. **Tag release** after merge
5. **Create GitHub release**

---

## Getting Help

- Check existing issues
- Review documentation
- Ask in pull request comments
- Open an issue with detailed description

---

## Documentation Standards

### Code Comments
```go
// Exported function must have comment
func DoSomething(x int) error {
    // Explain non-obvious logic
    if x < 0 {
        return errors.New("x must be positive")
    }
    return nil
}
```

### API Documentation
Update docs/api.md with:
- Endpoint description
- Request parameters
- Response examples
- Error cases
- Example cURL

### README Updates
Include:
- Feature overview
- Installation steps
- Usage examples
- Configuration options

---

## Performance Considerations

- Avoid N+1 queries
- Use database indexes
- Implement caching where appropriate
- Profile before optimizing
- Write benchmarks for critical paths

---

## Security Considerations

- Never hardcode secrets
- Validate all inputs
- Escape output when needed
- Use prepared statements for DB
- Keep dependencies updated
- Follow OWASP guidelines

---

## Common Issues

### Import Errors
```bash
make deps
go mod tidy
```

### Test Failures
```bash
# Run tests with verbose output
go test -v ./tests/unit/...

# Run specific test
go test -v -run TestName ./tests/unit/...
```

### Build Errors
```bash
# Clean build
make clean
make build
```

---

## Useful Commands

```bash
# Format and lint before commit
make fmt lint

# Run all tests
make test

# Build binary
make build

# View available commands
make help
```

---

## Questions?

- Check the documentation
- Review similar code in the repo
- Open an issue
- Ask in pull request

---

**Thank you for contributing!** 🎉

Your contributions make Todo App better for everyone.

---

**Last Updated:** January 20, 2026
