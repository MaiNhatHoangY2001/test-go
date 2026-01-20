package usecase

import (
	"context"
	"strings"
	"test-go/internal/features/todo/dto"
	"test-go/internal/features/todo/repository"
	errs "test-go/internal/shared/errors"
	"time"
)

type UpdateTodoUseCase struct {
	repository repository.TodoRepository
}

func NewUpdateTodoUseCase(repo repository.TodoRepository) *UpdateTodoUseCase {
	return &UpdateTodoUseCase{
		repository: repo,
	}
}

func (uc *UpdateTodoUseCase) Execute(ctx context.Context, input dto.UpdateTodoInput) (*dto.UpdateTodoOutput, error) {
	existingTodo, err := uc.repository.GetByID(ctx, input.ID)
	if err != nil {
		if strings.Contains(err.Error(), "not found") {
			return nil, errs.New(errs.NotFoundError, "Todo not found")
		}
		if strings.Contains(err.Error(), "invalid id") {
			return nil, errs.New(errs.BadRequestError, "Invalid todo ID format")
		}
		return nil, errs.Wrap(err, errs.DatabaseError, "Failed to retrieve todo")
	}

	existingTodo.Title = input.Title
	existingTodo.Description = input.Description
	existingTodo.Completed = input.Completed
	existingTodo.UpdatedAt = time.Now()

	if err := uc.repository.Update(ctx, existingTodo); err != nil {
		if strings.Contains(err.Error(), "not found") {
			return nil, errs.New(errs.NotFoundError, "Todo not found")
		}
		return nil, errs.Wrap(err, errs.DatabaseError, "Failed to update todo")
	}

	return &dto.UpdateTodoOutput{
		ID:          existingTodo.ID,
		Title:       existingTodo.Title,
		Description: existingTodo.Description,
		Completed:   existingTodo.Completed,
		CreatedAt:   existingTodo.CreatedAt,
		UpdatedAt:   existingTodo.UpdatedAt,
	}, nil
}
