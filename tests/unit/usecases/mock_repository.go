package usecases

import (
	"context"
	"errors"
	"sync"
	"test-go/internal/domain/entities"
)

// MockTodoRepository is a simple mock implementation for testing
type MockTodoRepository struct {
	mu    sync.RWMutex
	todos map[string]*entities.Todo
}

func NewMockTodoRepository() *MockTodoRepository {
	return &MockTodoRepository{
		todos: make(map[string]*entities.Todo),
	}
}

func (m *MockTodoRepository) Create(ctx context.Context, todo *entities.Todo) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.todos[todo.ID.Hex()] = todo
	return nil
}

func (m *MockTodoRepository) GetByID(ctx context.Context, userID, id string) (*entities.Todo, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	if todo, exists := m.todos[id]; exists {
		if todo.UserID == userID {
			return todo, nil
		}
	}
	return nil, errors.New("todo not found")
}

func (m *MockTodoRepository) GetAll(ctx context.Context, userID string, page, limit int) ([]*entities.Todo, int64, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	
	// Filter todos by user ID
	userTodos := make([]*entities.Todo, 0)
	for _, todo := range m.todos {
		if todo.UserID == userID {
			userTodos = append(userTodos, todo)
		}
	}
	
	totalCount := int64(len(userTodos))
	
	// Apply pagination using totalCount for boundary checking
	start := (page - 1) * limit
	end := start + limit
	if start > int(totalCount) {
		return []*entities.Todo{}, totalCount, nil
	}
	if end > int(totalCount) {
		end = int(totalCount)
	}
	
	return userTodos[start:end], totalCount, nil
}

func (m *MockTodoRepository) Update(ctx context.Context, userID string, todo *entities.Todo) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	key := todo.ID.Hex()
	existingTodo, exists := m.todos[key]
	if !exists || existingTodo.UserID != userID {
		return errors.New("todo not found")
	}
	m.todos[key] = todo
	return nil
}

func (m *MockTodoRepository) Delete(ctx context.Context, userID, id string) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	todo, exists := m.todos[id]
	if !exists || todo.UserID != userID {
		return errors.New("todo not found")
	}
	delete(m.todos, id)
	return nil
}
