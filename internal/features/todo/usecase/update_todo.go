package usecase

import (
	"context"
	"test-go/internal/domain/repositories"
	"test-go/internal/features/todo/dto"
	"time"
)

type UpdateTodoUseCase struct {
	repository repositories.TodoRepository
}

func NewUpdateTodoUseCase(repo repositories.TodoRepository) *UpdateTodoUseCase {
	return &UpdateTodoUseCase{
		repository: repo,
	}
}

func (uc *UpdateTodoUseCase) Execute(ctx context.Context, userID string, input dto.UpdateTodoInput) (*dto.UpdateTodoOutput, error) {
	existingTodo, err := uc.repository.GetByID(ctx, userID, input.ID)
	if err != nil {
		return nil, err
	}

	existingTodo.Title = input.Title
	existingTodo.Description = input.Description
	existingTodo.Completed = input.Completed
	existingTodo.UpdatedAt = time.Now()

	if err := uc.repository.Update(ctx, userID, existingTodo); err != nil {
		return nil, err
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
