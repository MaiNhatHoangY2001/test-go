package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"test-go/internal/domain/entities"
	"test-go/internal/features/auth/dto"
	"test-go/internal/features/auth/handler"
	"test-go/internal/features/auth/usecase"
	"test-go/pkg/logger"
	"test-go/tests/helpers"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AuthHandlerTestSuite struct {
	suite.Suite
	mockUserRepo *helpers.MockUserRepository
	authHandler  *handler.AuthHandler
	logger       *logrus.Logger
}

func (suite *AuthHandlerTestSuite) SetupTest() {
	gin.SetMode(gin.TestMode)
	suite.mockUserRepo = new(helpers.MockUserRepository)
	suite.logger = logger.InitLogger()

	// Create use cases with mock repo
	signUpUseCase := usecase.NewSignUpUseCase(suite.mockUserRepo)
	loginUseCase := usecase.NewLoginUseCase(suite.mockUserRepo)

	suite.authHandler = handler.NewAuthHandler(signUpUseCase, loginUseCase, suite.logger)
}

func (suite *AuthHandlerTestSuite) TestSignUp_Success() {
	// Arrange
	reqBody := dto.SignUpRequest{
		Email:    "test@example.com",
		Password: "TestPassword123",
		Name:     "Test User",
	}

	suite.mockUserRepo.On("FindByEmail", context.Background(), reqBody.Email).Return(nil, nil)
	suite.mockUserRepo.On("Create", context.Background(), mock.MatchedBy(func(u *entities.User) bool {
		return u.Email == reqBody.Email && u.Name == reqBody.Name
	})).Return(nil)

	body, _ := json.Marshal(reqBody)
	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/auth/signup", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	c, _ := gin.CreateTestContext(w)
	c.Request = req

	// Act
	suite.authHandler.SignUp(c)

	// Assert
	assert.Equal(suite.T(), http.StatusCreated, w.Code)
	suite.mockUserRepo.AssertExpectations(suite.T())
}

func (suite *AuthHandlerTestSuite) TestSignUp_InvalidEmail() {
	// Arrange
	reqBody := dto.SignUpRequest{
		Email:    "invalid-email",
		Password: "TestPassword123",
		Name:     "Test User",
	}

	body, _ := json.Marshal(reqBody)
	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/auth/signup", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	c, _ := gin.CreateTestContext(w)
	c.Request = req

	// Act
	suite.authHandler.SignUp(c)

	// Assert
	assert.Equal(suite.T(), http.StatusBadRequest, w.Code)
}

func (suite *AuthHandlerTestSuite) TestSignUp_WeakPassword() {
	// Arrange
	reqBody := dto.SignUpRequest{
		Email:    "test@example.com",
		Password: "weak",
		Name:     "Test User",
	}

	body, _ := json.Marshal(reqBody)
	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/auth/signup", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	c, _ := gin.CreateTestContext(w)
	c.Request = req

	// Act
	suite.authHandler.SignUp(c)

	// Assert
	assert.Equal(suite.T(), http.StatusBadRequest, w.Code)
}

func (suite *AuthHandlerTestSuite) TestSignUp_DuplicateEmail() {
	// Arrange
	reqBody := dto.SignUpRequest{
		Email:    "existing@example.com",
		Password: "TestPassword123",
		Name:     "Test User",
	}

	existingUser := helpers.CreateTestUser(
		primitive.NewObjectID().Hex(),
		reqBody.Email,
		"hashedPassword",
		"Existing User",
	)

	suite.mockUserRepo.On("FindByEmail", context.Background(), reqBody.Email).Return(existingUser, nil)

	body, _ := json.Marshal(reqBody)
	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/auth/signup", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	c, _ := gin.CreateTestContext(w)
	c.Request = req

	// Act
	suite.authHandler.SignUp(c)

	// Assert
	assert.Equal(suite.T(), http.StatusConflict, w.Code)
	suite.mockUserRepo.AssertExpectations(suite.T())
}

func (suite *AuthHandlerTestSuite) TestLogin_Success() {
	// Arrange
	reqBody := dto.LoginRequest{
		Email:    "test@example.com",
		Password: "TestPassword123",
	}

	user := helpers.CreateTestUser(
		primitive.NewObjectID().Hex(),
		reqBody.Email,
		"hashedPassword",
		"Test User",
	)

	suite.mockUserRepo.On("FindByEmail", context.Background(), reqBody.Email).Return(user, nil)

	body, _ := json.Marshal(reqBody)
	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/auth/login", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	c, _ := gin.CreateTestContext(w)
	c.Request = req

	// Act
	suite.authHandler.Login(c)

	// Assert
	assert.Equal(suite.T(), http.StatusOK, w.Code)
	suite.mockUserRepo.AssertExpectations(suite.T())
}

func (suite *AuthHandlerTestSuite) TestLogin_UserNotFound() {
	// Arrange
	reqBody := dto.LoginRequest{
		Email:    "nonexistent@example.com",
		Password: "TestPassword123",
	}

	suite.mockUserRepo.On("FindByEmail", context.Background(), reqBody.Email).Return(nil, nil)

	body, _ := json.Marshal(reqBody)
	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/auth/login", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	c, _ := gin.CreateTestContext(w)
	c.Request = req

	// Act
	suite.authHandler.Login(c)

	// Assert
	assert.Equal(suite.T(), http.StatusUnauthorized, w.Code)
	suite.mockUserRepo.AssertExpectations(suite.T())
}

func (suite *AuthHandlerTestSuite) TestLogin_InvalidEmail() {
	// Arrange
	reqBody := dto.LoginRequest{
		Email:    "invalid-email",
		Password: "TestPassword123",
	}

	body, _ := json.Marshal(reqBody)
	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/auth/login", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	c, _ := gin.CreateTestContext(w)
	c.Request = req

	// Act
	suite.authHandler.Login(c)

	// Assert
	assert.Equal(suite.T(), http.StatusBadRequest, w.Code)
}

func TestAuthHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(AuthHandlerTestSuite))
}
