package repository

import (
"context"
"test-go/internal/features/todo/entity"
)

type TodoRepository interface {
Create(ctx context.Context, todo *entity.Todo) error
GetByID(ctx context.Context, id string) (*entity.Todo, error)
GetAll(ctx context.Context) ([]*entity.Todo, error)
Update(ctx context.Context, todo *entity.Todo) error
Delete(ctx context.Context, id string) error
}
