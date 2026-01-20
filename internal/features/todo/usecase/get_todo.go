package usecase

import (
	"context"
	"strings"
	"test-go/internal/domain/repositories"
	"test-go/internal/features/todo/dto"
	errs "test-go/internal/shared/errors"
)

type GetTodoUseCase struct {
	repository repositories.TodoRepository
}

func NewGetTodoUseCase(repo repositories.TodoRepository) *GetTodoUseCase {
	return &GetTodoUseCase{
		repository: repo,
	}
}

func (uc *GetTodoUseCase) Execute(ctx context.Context, input dto.GetTodoInput) (*dto.GetTodoOutput, error) {
	todo, err := uc.repository.GetByID(ctx, input.ID)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return nil, errs.New(errs.NotFoundError, "Todo not found")
		}
		if strings.Contains(err.Error(), "invalid id") {
			return nil, errs.New(errs.BadRequestError, "Invalid todo ID format")
		}
		return nil, errs.Wrap(err, errs.DatabaseError, "Failed to retrieve todo")
	}

	return &dto.GetTodoOutput{
		ID:          todo.ID,
		Title:       todo.Title,
		Description: todo.Description,
		Completed:   todo.Completed,
		CreatedAt:   todo.CreatedAt,
		UpdatedAt:   todo.UpdatedAt,
	}, nil
}
