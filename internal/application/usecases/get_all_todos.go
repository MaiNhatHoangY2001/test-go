package usecases

import (
	"context"
	"test-go/internal/domain/repositories"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type GetAllTodosUseCase struct {
	repository repositories.TodoRepository
}

func NewGetAllTodosUseCase(repo repositories.TodoRepository) *GetAllTodosUseCase {
	return &GetAllTodosUseCase{
		repository: repo,
	}
}

type GetAllTodosOutput struct {
	ID          primitive.ObjectID `json:"id"`
	Title       string             `json:"title"`
	Description string             `json:"description"`
	Completed   bool               `json:"completed"`
	CreatedAt   time.Time          `json:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at"`
}

func (uc *GetAllTodosUseCase) Execute(ctx context.Context) ([]GetAllTodosOutput, error) {
	todos, err := uc.repository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	outputs := make([]GetAllTodosOutput, len(todos))
	for i, todo := range todos {
		outputs[i] = GetAllTodosOutput{
			ID:          todo.ID,
			Title:       todo.Title,
			Description: todo.Description,
			Completed:   todo.Completed,
			CreatedAt:   todo.CreatedAt,
			UpdatedAt:   todo.UpdatedAt,
		}
	}

	return outputs, nil
}
