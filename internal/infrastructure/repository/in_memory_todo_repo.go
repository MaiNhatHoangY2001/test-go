package repository

import (
	"context"
	"errors"
	"sync"
	"test-go/internal/domain/entities"
	"test-go/internal/domain/repositories"
)

type InMemoryTodoRepository struct {
	mu    sync.RWMutex
	todos map[string]*entities.Todo
}

func NewInMemoryTodoRepository() repositories.TodoRepository {
	return &InMemoryTodoRepository{
		todos: make(map[string]*entities.Todo),
	}
}

func (i *InMemoryTodoRepository) Create(ctx context.Context, todo *entities.Todo) error {
	i.mu.Lock()
	defer i.mu.Unlock()

	if _, exists := i.todos[todo.ID]; exists {
		return errors.New("todo already exists")
	}

	i.todos[todo.ID] = todo
	return nil
}

func (i *InMemoryTodoRepository) Delete(ctx context.Context, id string) error {
	i.mu.Lock()
	defer i.mu.Unlock()

	if _, exists := i.todos[id]; !exists {
		return errors.New("todo not found")
	}

	delete(i.todos, id)
	return nil
}

func (i *InMemoryTodoRepository) GetAll(ctx context.Context) ([]*entities.Todo, error) {
	i.mu.RLock()
	defer i.mu.RUnlock()

	todos := make([]*entities.Todo, 0, len(i.todos))
	for _, todo := range i.todos {
		todos = append(todos, todo)
	}

	return todos, nil
}

func (i *InMemoryTodoRepository) GetByID(ctx context.Context, id string) (*entities.Todo, error) {
	i.mu.RLock()
	defer i.mu.RUnlock()

	todo, exists := i.todos[id]
	if !exists {
		return nil, errors.New("todo not found")
	}

	return todo, nil
}

func (i *InMemoryTodoRepository) Update(ctx context.Context, todo *entities.Todo) error {
	i.mu.Lock()
	defer i.mu.Unlock()

	if _, exists := i.todos[todo.ID]; !exists {
		return errors.New("todo not found")
	}

	i.todos[todo.ID] = todo
	return nil
}
