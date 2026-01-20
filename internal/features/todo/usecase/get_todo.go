package usecase

import (
	"context"
	"test-go/internal/domain/repositories"
	"test-go/internal/features/todo/dto"
)

type GetTodoUseCase struct {
	repository repositories.TodoRepository
}

func NewGetTodoUseCase(repo repositories.TodoRepository) *GetTodoUseCase {
	return &GetTodoUseCase{
		repository: repo,
	}
}

func (uc *GetTodoUseCase) Execute(ctx context.Context, userID string, input dto.GetTodoInput) (*dto.GetTodoOutput, error) {
	todo, err := uc.repository.GetByID(ctx, userID, input.ID)
	if err != nil {
		return nil, err
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
