package usecase

import (
	"context"
	"strings"
	"test-go/internal/domain/repositories"
	"test-go/internal/features/todo/dto"
	errs "test-go/internal/shared/errors"
)

type DeleteTodoUseCase struct {
	repository repositories.TodoRepository
}

func NewDeleteTodoUseCase(repo repositories.TodoRepository) *DeleteTodoUseCase {
	return &DeleteTodoUseCase{
		repository: repo,
	}
}

func (uc *DeleteTodoUseCase) Execute(ctx context.Context, input dto.DeleteTodoInput) error {
	err := uc.repository.Delete(ctx, input.ID)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return errs.New(errs.NotFoundError, "Todo not found")
		}
		if strings.Contains(err.Error(), "invalid id") {
			return errs.New(errs.BadRequestError, "Invalid todo ID format")
		}
		return errs.Wrap(err, errs.DatabaseError, "Failed to delete todo")
	}

	return nil
}
