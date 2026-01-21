# Configuration Guide

## Environment Variables

Create a `.env` file in the project root with the following variables:

### Server Configuration
```bash
# Server port
PORT=8080

# Environment (development, staging, production)
NODE_ENV=development
```

### Database Configuration
```bash
# MongoDB connection string
MONGO_URI=mongodb://localhost:27017

# Database name
DATABASE_NAME=todos

# Collection name for todos
COLLECTION_NAME=todos
```

### Security Configuration
```bash
# JWT Secret Key - change this in production!
# Generate a secure key: openssl rand -base64 32
JWT_SECRET=your-super-secret-key-here-change-in-production

# Token expiration (in hours)
JWT_EXPIRATION=24
```

### Logging Configuration
```bash
# Log level (debug, info, warn, error)
LOG_LEVEL=info

# Log format (json, text)
LOG_FORMAT=json
```

---

## Complete .env.example

```bash
# Server Configuration
PORT=8080
NODE_ENV=development

# Database Configuration
MONGO_URI=mongodb://localhost:27017
DATABASE_NAME=todos
COLLECTION_NAME=todos

# Security
JWT_SECRET=your-secret-key-change-in-production
JWT_EXPIRATION=24

# Logging
LOG_LEVEL=info
LOG_FORMAT=json
```

---

## Loading Configuration

The application loads configuration from:
1. Environment variables (highest priority)
2. `.env` file
3. Default values (lowest priority)

### Precedence Example

```go
// If environment variable exists, use it
PORT := GetEnv("PORT", "8080")  // Uses env var if PORT is set
                                  // Otherwise uses "8080"
```

---

## MongoDB Setup

### Local Development with Docker

```bash
# Start MongoDB container
docker run -d \
  -p 27017:27017 \
  --name mongo \
  -e MONGO_INITDB_ROOT_USERNAME=root \
  -e MONGO_INITDB_ROOT_PASSWORD=password \
  mongo:latest

# Connect using:
# MONGO_URI=mongodb://root:password@localhost:27017
```

### Using Docker Compose

```bash
# Start MongoDB and application
docker-compose up -d

# View logs
docker-compose logs -f mongo

# Stop
docker-compose down
```

---

## Running in Different Environments

### Development
```bash
# Create .env for development
PORT=8080
NODE_ENV=development
MONGO_URI=mongodb://localhost:27017
DATABASE_NAME=todos_dev
JWT_SECRET=dev-secret-key

# Run
make run
```

### Staging
```bash
# Set environment variables for staging
export PORT=8080
export NODE_ENV=staging
export MONGO_URI=mongodb://staging-db:27017
export DATABASE_NAME=todos_staging
export JWT_SECRET=$(openssl rand -base64 32)

# Run
make build && ./server
```

### Production
```bash
# Set production environment variables
# Recommended: Use container orchestration (Kubernetes, Docker Swarm)
export PORT=8080
export NODE_ENV=production
export MONGO_URI=mongodb://prod-db-cluster
export DATABASE_NAME=todos
export JWT_SECRET=$(openssl rand -base64 32)
export LOG_LEVEL=warn

# Run
./server
```

---

## Docker Compose Configuration

The `docker-compose.yml` file includes:

```yaml
version: '3.8'

services:
  mongo:
    image: mongo:latest
    ports:
      - "27017:27017"
    environment:
      MONGO_INITDB_ROOT_USERNAME: root
      MONGO_INITDB_ROOT_PASSWORD: password
    volumes:
      - mongo_data:/data/db

  app:
    build: .
    ports:
      - "8080:8080"
    environment:
      MONGO_URI: mongodb://root:password@mongo:27017
      DATABASE_NAME: todos
      JWT_SECRET: docker-secret-key
    depends_on:
      - mongo

volumes:
  mongo_data:
```

---

## Database Indexes

Recommended indexes for performance:

```javascript
// Create indexes in MongoDB
db.todos.createIndex({ "user_id": 1 })
db.todos.createIndex({ "created_at": -1 })
db.users.createIndex({ "email": 1 }, { unique: true })
```

---

## Security Recommendations

### JWT Secret
- **Development**: Any string is fine
- **Production**: Use a strong random string
  ```bash
  # Generate secure secret
  openssl rand -base64 32
  ```

### MongoDB
- **Development**: No authentication required
- **Production**: Enable authentication
  ```
  MONGO_URI=mongodb://username:password@host:port/database
  ```

### Environment Variables
- Never commit `.env` file to git
- Use `.env.example` for documentation
- Rotate secrets regularly in production

---

## Troubleshooting

### "Connection refused" error
- Check if MongoDB is running
- Verify MONGO_URI is correct
- Check firewall/network settings

### "Invalid token" error
- Ensure JWT_SECRET matches across restarts
- Check token expiration time
- Verify Authorization header format

### Port already in use
- Change PORT environment variable
- Or kill process using the port:
  ```bash
  # Linux/Mac
  lsof -i :8080 | grep -v PID | awk '{print $2}' | xargs kill

  # Windows
  netstat -ano | findstr :8080
  taskkill /PID {PID} /F
  ```

---

**Last Updated:** January 20, 2026
