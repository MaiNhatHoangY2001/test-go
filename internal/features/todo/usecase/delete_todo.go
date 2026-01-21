package usecase

import (
	"context"
	"test-go/internal/domain/repositories"
	"test-go/internal/features/todo/dto"
)

type DeleteTodoUseCase struct {
	repository repositories.TodoRepository
}

func NewDeleteTodoUseCase(repo repositories.TodoRepository) *DeleteTodoUseCase {
	return &DeleteTodoUseCase{
		repository: repo,
	}
}

func (uc *DeleteTodoUseCase) Execute(ctx context.Context, userID string, input dto.DeleteTodoInput) error {
	return uc.repository.Delete(ctx, userID, input.ID)
}
