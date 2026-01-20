package repositories

import (
	"context"
	"test-go/internal/domain/entities"
)

type TodoRepository interface {
	Create(ctx context.Context, todo *entities.Todo) error
	GetByID(ctx context.Context, id string) (*entities.Todo, error)
	GetAll(ctx context.Context, page, limit int) ([]*entities.Todo, int64, error)
	Update(ctx context.Context, todo *entities.Todo) error
	Delete(ctx context.Context, id string) error
}
