package usecase

import (
	"context"
	"test-go/internal/domain/entities"
	"test-go/internal/features/todo/dto"
	"test-go/internal/features/todo/repository"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateTodoUseCase struct {
	repository repository.TodoRepository
}

func NewCreateTodoUseCase(repo repository.TodoRepository) *CreateTodoUseCase {
	return &CreateTodoUseCase{
		repository: repo,
	}
}

func (uc *CreateTodoUseCase) Execute(ctx context.Context, input dto.CreateTodoInput) (*dto.CreateTodoOutput, error) {
	todo := &entities.Todo{
		ID:          primitive.NewObjectID(),
		Title:       input.Title,
		Description: input.Description,
		Completed:   false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if err := uc.repository.Create(ctx, todo); err != nil {
		return nil, err
	}

	return &dto.CreateTodoOutput{
		ID:          todo.ID,
		Title:       todo.Title,
		Description: todo.Description,
		Completed:   todo.Completed,
		CreatedAt:   todo.CreatedAt,
		UpdatedAt:   todo.UpdatedAt,
	}, nil
}
