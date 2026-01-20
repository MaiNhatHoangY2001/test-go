package usecases

import (
	"context"
	"test-go/internal/domain/repositories"
	"time"
)

type GetTodoUseCase struct {
	repository repositories.TodoRepository
}

func NewGetTodoUseCase(repo repositories.TodoRepository) *GetTodoUseCase {
	return &GetTodoUseCase{
		repository: repo,
	}
}

type GetTodoInput struct {
	ID string
}

type GetTodoOutput struct {
	ID          string
	Title       string
	Description string
	Completed   bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (uc *GetTodoUseCase) Execute(ctx context.Context, input GetTodoInput) (*GetTodoOutput, error) {
	todo, err := uc.repository.GetByID(ctx, input.ID)
	if err != nil {
		return nil, err
	}

	return &GetTodoOutput{
		ID:          todo.ID,
		Title:       todo.Title,
		Description: todo.Description,
		Completed:   todo.Completed,
		CreatedAt:   todo.CreatedAt,
		UpdatedAt:   todo.UpdatedAt,
	}, nil
}
