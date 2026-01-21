package integration

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	"test-go/internal/features/auth/dto"
	"test-go/internal/infrastructure/database/mongodb"
	"test-go/internal/shared/config"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/mongo"
)

type APITestSuite struct {
	suite.Suite
	app    *config.App
	client *mongo.Client
}

func (suite *APITestSuite) SetupSuite() {
	// Load environment variables from .env file
	_ = godotenv.Load("../../.env")

	mongoURI := os.Getenv("MONGO_URI")
	if mongoURI == "" {
		mongoURI = "mongodb://localhost:27017"
	}

	client, err := mongodb.NewMongoClient(context.Background(), mongoURI)
	require.NoError(suite.T(), err)
	suite.client = client

	appConfig := &config.AppConfig{
		MongoURI:       mongoURI,
		DatabaseName:   "test-go",
		CollectionName: "todos",
	}

	app, err := config.NewApp(appConfig)
	require.NoError(suite.T(), err)
	suite.app = app
}

func (suite *APITestSuite) TearDownSuite() {
	if suite.client != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		mongodb.CloseMongoClient(ctx, suite.client)
	}
}

func (suite *APITestSuite) CleanupCollections() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	db := suite.client.Database("test-go")
	db.Collection("todos").DeleteMany(ctx, map[string]interface{}{})
	db.Collection("users").DeleteMany(ctx, map[string]interface{}{})
}

// Helper method to create a user and get authentication token
func (suite *APITestSuite) GetAuthToken(email, password, name string) string {
	// Sign up
	signUpReq := dto.SignUpRequest{
		Email:    email,
		Password: password,
		Name:     name,
	}
	body, _ := json.Marshal(signUpReq)
	request := httptest.NewRequest(http.MethodPost, "/api/v1/auth/signup", bytes.NewBuffer(body))
	request.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	suite.app.Router.ServeHTTP(w, request)

	// Login
	loginReq := dto.LoginRequest{
		Email:    email,
		Password: password,
	}
	body, _ = json.Marshal(loginReq)
	request = httptest.NewRequest(http.MethodPost, "/api/v1/auth/login", bytes.NewBuffer(body))
	request.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	suite.app.Router.ServeHTTP(w, request)

	var response map[string]interface{}
	json.Unmarshal(w.Body.Bytes(), &response)

	// Extract token from response
	if data, ok := response["data"].(map[string]interface{}); ok {
		if token, ok := data["token"].(string); ok {
			return token
		}
	}
	return ""
}

func (suite *APITestSuite) TestSignUpSuccess() {
	suite.CleanupCollections()

	req := dto.SignUpRequest{
		Email:    "test@example.com",
		Password: "SecurePassword123",
		Name:     "Test User",
	}

	body, _ := json.Marshal(req)
	request := httptest.NewRequest(http.MethodPost, "/api/v1/auth/signup", bytes.NewBuffer(body))
	request.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	suite.app.Router.ServeHTTP(w, request)

	assert.Equal(suite.T(), http.StatusCreated, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), response["data"])
}

func (suite *APITestSuite) TestSignUpInvalidEmail() {
	suite.CleanupCollections()

	req := dto.SignUpRequest{
		Email:    "invalid-email",
		Password: "SecurePassword123",
		Name:     "Test User",
	}

	body, _ := json.Marshal(req)
	request := httptest.NewRequest(http.MethodPost, "/api/v1/auth/signup", bytes.NewBuffer(body))
	request.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	suite.app.Router.ServeHTTP(w, request)

	assert.Equal(suite.T(), http.StatusBadRequest, w.Code)
}

func (suite *APITestSuite) TestSignUpWeakPassword() {
	suite.CleanupCollections()

	req := dto.SignUpRequest{
		Email:    "test@example.com",
		Password: "weak",
		Name:     "Test User",
	}

	body, _ := json.Marshal(req)
	request := httptest.NewRequest(http.MethodPost, "/api/v1/auth/signup", bytes.NewBuffer(body))
	request.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	suite.app.Router.ServeHTTP(w, request)

	assert.Equal(suite.T(), http.StatusBadRequest, w.Code)
}

func (suite *APITestSuite) TestLoginSuccess() {
	suite.CleanupCollections()

	signUpReq := dto.SignUpRequest{
		Email:    "login@example.com",
		Password: "SecurePassword123",
		Name:     "Login User",
	}
	body, _ := json.Marshal(signUpReq)
	request := httptest.NewRequest(http.MethodPost, "/api/v1/auth/signup", bytes.NewBuffer(body))
	request.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	suite.app.Router.ServeHTTP(w, request)

	loginReq := dto.LoginRequest{
		Email:    "login@example.com",
		Password: "SecurePassword123",
	}
	body, _ = json.Marshal(loginReq)
	request = httptest.NewRequest(http.MethodPost, "/api/v1/auth/login", bytes.NewBuffer(body))
	request.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()

	suite.app.Router.ServeHTTP(w, request)

	assert.Equal(suite.T(), http.StatusOK, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(suite.T(), err)
	assert.NotNil(suite.T(), response["data"])
}

func (suite *APITestSuite) TestLoginInvalidCredentials() {
	suite.CleanupCollections()

	loginReq := dto.LoginRequest{
		Email:    "nonexistent@example.com",
		Password: "WrongPassword123",
	}
	body, _ := json.Marshal(loginReq)
	request := httptest.NewRequest(http.MethodPost, "/api/v1/auth/login", bytes.NewBuffer(body))
	request.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()

	suite.app.Router.ServeHTTP(w, request)

	assert.Equal(suite.T(), http.StatusBadRequest, w.Code)
}

func (suite *APITestSuite) TestRequestIDMiddleware() {
	request := httptest.NewRequest(http.MethodGet, "/api/v1/todos", nil)
	w := httptest.NewRecorder()

	suite.app.Router.ServeHTTP(w, request)

	requestID := w.Header().Get("X-Request-ID")
	assert.NotEmpty(suite.T(), requestID)
}

func TestAPITestSuite(t *testing.T) {
	suite.Run(t, new(APITestSuite))
}
