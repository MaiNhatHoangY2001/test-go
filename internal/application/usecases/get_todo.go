package usecases

import (
	"context"
	"test-go/internal/domain/repositories"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
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
	ID          primitive.ObjectID `json:"id"`
	Title       string             `json:"title"`
	Description string             `json:"description"`
	Completed   bool               `json:"completed"`
	CreatedAt   time.Time          `json:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at"`
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
