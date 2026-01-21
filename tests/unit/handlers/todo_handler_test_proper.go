package handlers

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"test-go/internal/domain/entities"
	"test-go/internal/features/todo/dto"
	"test-go/internal/features/todo/handler"
	"test-go/internal/features/todo/usecase"
	"test-go/pkg/logger"
	"test-go/tests/helpers"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TodoHandlerTestSuite struct {
	suite.Suite
	mockTodoRepo *helpers.MockTodoRepository
	todoHandler  *handler.TodoHandler
	logger       *logrus.Logger
}

func (suite *TodoHandlerTestSuite) SetupTest() {
	gin.SetMode(gin.TestMode)
	suite.mockTodoRepo = new(helpers.MockTodoRepository)
	suite.logger = logger.InitLogger()

	createUC := usecase.NewCreateTodoUseCase(suite.mockTodoRepo)
	getTodoUC := usecase.NewGetTodoUseCase(suite.mockTodoRepo)
	getAllUC := usecase.NewGetAllTodosUseCase(suite.mockTodoRepo)
	updateUC := usecase.NewUpdateTodoUseCase(suite.mockTodoRepo)
	deleteUC := usecase.NewDeleteTodoUseCase(suite.mockTodoRepo)

	suite.todoHandler = handler.NewTodoHandler(createUC, getTodoUC, getAllUC, updateUC, deleteUC, suite.logger)
}

func (suite *TodoHandlerTestSuite) TestCreateTodo_Success() {
	// Arrange
	reqBody := dto.CreateTodoInput{
		Title:       "Test Todo",
		Description: "Test Description",
	}

	suite.mockTodoRepo.On("Create", context.Background(), mock.MatchedBy(func(t *entities.Todo) bool {
		return t.Title == reqBody.Title
	})).Return(nil)

	body, _ := json.Marshal(reqBody)
	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/todos", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Request-ID", "test-123")

	c, _ := gin.CreateTestContext(w)
	c.Request = req

	// Act
	suite.todoHandler.CreateTodo(c)

	// Assert
	assert.Equal(suite.T(), http.StatusCreated, w.Code)
	suite.mockTodoRepo.AssertExpectations(suite.T())
}

func (suite *TodoHandlerTestSuite) TestCreateTodo_InvalidInput() {
	// Arrange
	reqBody := dto.CreateTodoInput{
		Title:       "",
		Description: "",
	}

	body, _ := json.Marshal(reqBody)
	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/todos", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Request-ID", "test-123")

	c, _ := gin.CreateTestContext(w)
	c.Request = req

	// Act
	suite.todoHandler.CreateTodo(c)

	// Assert
	assert.Equal(suite.T(), http.StatusBadRequest, w.Code)
}

func (suite *TodoHandlerTestSuite) TestGetAllTodos_Success() {
	// Arrange
	todos := []*entities.Todo{
		helpers.CreateTestTodo(primitive.NewObjectID().Hex(), "Todo 1", "Desc 1", false),
		helpers.CreateTestTodo(primitive.NewObjectID().Hex(), "Todo 2", "Desc 2", true),
	}

	suite.mockTodoRepo.On("GetAll", context.Background()).Return(todos, nil)

	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/todos", nil)
	req.Header.Set("X-Request-ID", "test-123")

	c, _ := gin.CreateTestContext(w)
	c.Request = req

	// Act
	suite.todoHandler.GetAllTodos(c)

	// Assert
	assert.Equal(suite.T(), http.StatusOK, w.Code)
	suite.mockTodoRepo.AssertExpectations(suite.T())
}

func (suite *TodoHandlerTestSuite) TestGetTodo_Success() {
	// Arrange
	todoID := primitive.NewObjectID().Hex()
	todo := helpers.CreateTestTodo(todoID, "Test Todo", "Test Desc", false)

	suite.mockTodoRepo.On("GetByID", context.Background(), todoID).Return(todo, nil)

	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/todos/"+todoID, nil)
	req.Header.Set("X-Request-ID", "test-123")

	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "id", Value: todoID}}
	c.Request = req

	// Act
	suite.todoHandler.GetTodo(c)

	// Assert
	assert.Equal(suite.T(), http.StatusOK, w.Code)
	suite.mockTodoRepo.AssertExpectations(suite.T())
}

func (suite *TodoHandlerTestSuite) TestGetTodo_NotFound() {
	// Arrange
	todoID := primitive.NewObjectID().Hex()
	suite.mockTodoRepo.On("GetByID", context.Background(), todoID).Return(nil, nil)

	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/todos/"+todoID, nil)
	req.Header.Set("X-Request-ID", "test-123")

	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "id", Value: todoID}}
	c.Request = req

	// Act
	suite.todoHandler.GetTodo(c)

	// Assert
	assert.Equal(suite.T(), http.StatusNotFound, w.Code)
}

func (suite *TodoHandlerTestSuite) TestUpdateTodo_Success() {
	// Arrange
	todoID := primitive.NewObjectID().Hex()
	reqBody := dto.UpdateTodoInput{
		Title:       "Updated Title",
		Description: "Updated Desc",
		Completed:   true,
	}

	suite.mockTodoRepo.On("Update", context.Background(), mock.MatchedBy(func(t *entities.Todo) bool {
		return t.Title == reqBody.Title
	})).Return(nil)

	body, _ := json.Marshal(reqBody)
	w := httptest.NewRecorder()
	req := httptest.NewRequest("PUT", "/todos/"+todoID, bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Request-ID", "test-123")

	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "id", Value: todoID}}
	c.Request = req

	// Act
	suite.todoHandler.UpdateTodo(c)

	// Assert
	assert.Equal(suite.T(), http.StatusOK, w.Code)
}

func (suite *TodoHandlerTestSuite) TestDeleteTodo_Success() {
	// Arrange
	todoID := primitive.NewObjectID().Hex()
	suite.mockTodoRepo.On("Delete", context.Background(), todoID).Return(nil)

	w := httptest.NewRecorder()
	req := httptest.NewRequest("DELETE", "/todos/"+todoID, nil)
	req.Header.Set("X-Request-ID", "test-123")

	c, _ := gin.CreateTestContext(w)
	c.Params = gin.Params{{Key: "id", Value: todoID}}
	c.Request = req

	// Act
	suite.todoHandler.DeleteTodo(c)

	// Assert
	assert.Equal(suite.T(), http.StatusNoContent, w.Code)
	suite.mockTodoRepo.AssertExpectations(suite.T())
}

func TestTodoHandlerTestSuite(t *testing.T) {
	suite.Run(t, new(TodoHandlerTestSuite))
}
