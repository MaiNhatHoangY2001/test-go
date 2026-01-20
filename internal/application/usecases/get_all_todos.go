package usecases

import (
	"context"
	"test-go/internal/domain/repositories"
	"time"
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
	ID          string
	Title       string
	Description string
	Completed   bool
	CreatedAt   time.Time
	UpdatedAt   time.Time
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
