package usecases

import (
	"context"
	"test-go/internal/domain/entities"
	"test-go/internal/domain/repositories"
	"time"

	"github.com/google/uuid"
)

type CreateTodoUseCase struct {
	repository repositories.TodoRepository
}

func NewCreateTodoUseCase(repo repositories.TodoRepository) *CreateTodoUseCase {
	return &CreateTodoUseCase{
		repository: repo,
	}
}

type CreateTodoInput struct {
	Title       string `json:"title" binding:"required,min=1,max=200"`
	Description string `json:"description" binding:"max=1000"`
}

type CreateTodoOutput struct {
	ID          string
	Title       string
	Description string
	Completed   bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func (uc *CreateTodoUseCase) Execute(ctx context.Context, input CreateTodoInput) (*CreateTodoOutput, error) {
	todo := &entities.Todo{
		ID:          uuid.New().String(),
		Title:       input.Title,
		Description: input.Description,
		Completed:   false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := uc.repository.Create(ctx, todo); err != nil {
		return nil, err
	}

	return &CreateTodoOutput{
		ID:          todo.ID,
		Title:       todo.Title,
		Description: todo.Description,
		Completed:   todo.Completed,
		CreatedAt:   todo.CreatedAt,
		UpdatedAt:   todo.UpdatedAt,
	}, nil
}
