package helpers

import (
	"context"
	"test-go/internal/domain/entities"

	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// MockTodoRepository is a mock implementation of TodoRepository
type MockTodoRepository struct {
	mock.Mock
}

func (m *MockTodoRepository) Create(ctx context.Context, todo *entities.Todo) error {
	args := m.Called(ctx, todo)
	return args.Error(0)
}

func (m *MockTodoRepository) GetByID(ctx context.Context, id string) (*entities.Todo, error) {
	args := m.Called(ctx, id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.Todo), args.Error(1)
}

func (m *MockTodoRepository) GetAll(ctx context.Context, page, limit int) ([]*entities.Todo, int64, error) {
	args := m.Called(ctx, page, limit)
	if args.Get(0) == nil {
		return nil, 0, args.Error(2)
	}
	return args.Get(0).([]*entities.Todo), args.Get(1).(int64), args.Error(2)
}

func (m *MockTodoRepository) Update(ctx context.Context, todo *entities.Todo) error {
	args := m.Called(ctx, todo)
	return args.Error(0)
}

func (m *MockTodoRepository) Delete(ctx context.Context, id string) error {
	args := m.Called(ctx, id)
	return args.Error(0)
}

// MockUserRepository is a mock implementation of UserRepository
type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) Create(ctx context.Context, user *entities.User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

func (m *MockUserRepository) FindByEmail(ctx context.Context, email string) (*entities.User, error) {
	args := m.Called(ctx, email)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entities.User), args.Error(1)
}

// Helper to create a test todo
func CreateTestTodo(id, title, description string, completed bool) *entities.Todo {
	objID, _ := primitive.ObjectIDFromHex(id)
	return &entities.Todo{
		ID:          objID,
		Title:       title,
		Description: description,
		Completed:   completed,
	}
}

// Helper to create a test user
func CreateTestUser(id, email, password, name string) *entities.User {
	objID, _ := primitive.ObjectIDFromHex(id)
	return &entities.User{
		ID:       objID,
		Email:    email,
		Password: password,
		Name:     name,
	}
}
