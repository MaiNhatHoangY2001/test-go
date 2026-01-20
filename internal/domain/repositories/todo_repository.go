package repositories

import (
	"context"
	"test-go/internal/domain/entities"
)

type TodoRepository interface {
	Create(ctx context.Context, todo *entities.Todo) error
	GetByID(ctx context.Context, userID, id string) (*entities.Todo, error)
	GetAll(ctx context.Context, userID string, page, limit int) ([]*entities.Todo, int64, error)
	Update(ctx context.Context, userID string, todo *entities.Todo) error
	Delete(ctx context.Context, userID, id string) error
}
