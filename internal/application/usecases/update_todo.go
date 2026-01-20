package usecases

import (
	"context"
	"test-go/internal/domain/repositories"
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

type UpdateTodoInput struct {
	ID          string
	Title       string
	Description string
	Completed   bool
}

type UpdateTodoOutput struct {
	ID          string
	Title       string
	Description string
	Completed   bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (uc *UpdateTodoUseCase) Execute(ctx context.Context, input UpdateTodoInput) (*UpdateTodoOutput, error) {
	existingTodo, err := uc.repository.GetByID(ctx, input.ID)
	if err != nil {
		return nil, err
	}

	existingTodo.Title = input.Title
	existingTodo.Description = input.Description
	existingTodo.Completed = input.Completed
	existingTodo.UpdatedAt = time.Now()

	if err := uc.repository.Update(ctx, existingTodo); err != nil {
		return nil, err
	}

	return &UpdateTodoOutput{
		ID:          existingTodo.ID,
		Title:       existingTodo.Title,
		Description: existingTodo.Description,
		Completed:   existingTodo.Completed,
		CreatedAt:   existingTodo.CreatedAt,
		UpdatedAt:   existingTodo.UpdatedAt,
	}, nil
}
