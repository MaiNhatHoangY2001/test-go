package usecases

import (
	"context"
	"test-go/internal/domain/repositories"
)

type DeleteTodoUseCase struct {
	repository repositories.TodoRepository
}

func NewDeleteTodoUseCase(repo repositories.TodoRepository) *DeleteTodoUseCase {
	return &DeleteTodoUseCase{
		repository: repo,
	}
}

type DeleteTodoInput struct {
	ID string
}

func (uc *DeleteTodoUseCase) Execute(ctx context.Context, input DeleteTodoInput) error {
	return uc.repository.Delete(ctx, input.ID)
}
