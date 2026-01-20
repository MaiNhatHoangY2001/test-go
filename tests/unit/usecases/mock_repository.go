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

func (m *MockTodoRepository) GetByID(ctx context.Context, id string) (*entities.Todo, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	if todo, exists := m.todos[id]; exists {
		return todo, nil
	}
	return nil, errors.New("todo not found")
}

func (m *MockTodoRepository) GetAll(ctx context.Context) ([]*entities.Todo, error) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	todos := make([]*entities.Todo, 0, len(m.todos))
	for _, todo := range m.todos {
		todos = append(todos, todo)
	}
	return todos, nil
}

func (m *MockTodoRepository) Update(ctx context.Context, todo *entities.Todo) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.todos[todo.ID.Hex()] = todo
	return nil
}

func (m *MockTodoRepository) Delete(ctx context.Context, id string) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	delete(m.todos, id)
	return nil
}
