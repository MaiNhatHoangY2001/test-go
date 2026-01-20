# Todo App - API Documentation

## Base URL
```
http://localhost:8080/api/v1
```

## Authentication
Most endpoints require JWT Bearer token authentication:
```
Authorization: Bearer {token}
```

---

## Authentication Endpoints

### 1. Sign Up (Register New User)
Create a new user account.

**Endpoint:** `POST /auth/signup`

**Request Body:**
```json
{
  "email": "user@example.com",
  "password": "SecurePassword123",
  "name": "John Doe"
}
```

**Request Parameters:**
| Field | Type | Required | Validation |
|-------|------|----------|-----------|
| email | string | Yes | Valid email format, max 254 characters |
| password | string | Yes | 6-50 characters |
| name | string | Yes | 1-100 characters |

**Success Response (201 Created):**
```json
{
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "email": "user@example.com",
    "name": "John Doe"
  }
}
```

**Error Responses:**
- `400 Bad Request` - Invalid request format or validation error
- `409 Conflict` - Email already registered

---

### 2. Login
Authenticate user and receive JWT token.

**Endpoint:** `POST /auth/login`

**Request Body:**
```json
{
  "email": "user@example.com",
  "password": "SecurePassword123"
}
```

**Success Response (200 OK):**
```json
{
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "email": "user@example.com",
    "name": "John Doe"
  }
}
```

---

## Todo Endpoints

All todo endpoints require authentication.

### 3. Create Todo
**Endpoint:** `POST /todos`

**Request Body:**
```json
{
  "title": "Learn Go Programming",
  "description": "Master Go fundamentals",
  "completed": false
}
```

**Success Response (201 Created):**
```json
{
  "data": {
    "id": "507f1f77bcf86cd799439011",
    "title": "Learn Go Programming",
    "description": "Master Go fundamentals",
    "completed": false,
    "created_at": "2024-01-20T12:00:00Z",
    "updated_at": "2024-01-20T12:00:00Z"
  }
}
```

---

### 4. Get All Todos
**Endpoint:** `GET /todos`

**Success Response (200 OK):**
```json
{
  "data": [
    {
      "id": "507f1f77bcf86cd799439011",
      "title": "Learn Go Programming",
      "description": "Master Go fundamentals",
      "completed": false,
      "created_at": "2024-01-20T12:00:00Z",
      "updated_at": "2024-01-20T12:00:00Z"
    }
  ]
}
```

---

### 5. Get Todo by ID
**Endpoint:** `GET /todos/:id`

**Success Response (200 OK):**
```json
{
  "data": {
    "id": "507f1f77bcf86cd799439011",
    "title": "Learn Go Programming",
    "description": "Master Go fundamentals",
    "completed": false,
    "created_at": "2024-01-20T12:00:00Z",
    "updated_at": "2024-01-20T12:00:00Z"
  }
}
```

---

### 6. Update Todo
**Endpoint:** `PUT /todos/:id`

**Request Body:**
```json
{
  "title": "Learn Go Advanced Topics",
  "description": "Deep dive into concurrency",
  "completed": true
}
```

**Success Response (200 OK):**
```json
{
  "data": {
    "id": "507f1f77bcf86cd799439011",
    "title": "Learn Go Advanced Topics",
    "description": "Deep dive into concurrency",
    "completed": true,
    "created_at": "2024-01-20T12:00:00Z",
    "updated_at": "2024-01-20T15:30:00Z"
  }
}
```

---

### 7. Delete Todo
**Endpoint:** `DELETE /todos/:id`

**Success Response (204 No Content)**

---

## Health Check

**Endpoint:** `GET /health` (No authentication required)

**Success Response (200 OK):**
```json
{
  "status": "healthy",
  "timestamp": 1705761600
}
```

---

## Error Responses

**Format:**
```json
{
  "code": "ERROR_CODE",
  "message": "Error message",
  "details": "Additional details (optional)"
}
```

**Error Codes:**
- `VALIDATION_ERROR` (400)
- `BAD_REQUEST` (400)
- `UNAUTHORIZED` (401)
- `FORBIDDEN` (403)
- `NOT_FOUND` (404)
- `CONFLICT` (409)
- `INTERNAL_ERROR` (500)

---

**Last Updated:** January 20, 2026
**API Version:** 1.0.0
