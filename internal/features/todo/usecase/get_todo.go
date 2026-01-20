package usecase

import (
"context"
"test-go/internal/features/todo/dto"
"test-go/internal/features/todo/repository"
)

type GetTodoUseCase struct {
repository repository.TodoRepository
}

func NewGetTodoUseCase(repo repository.TodoRepository) *GetTodoUseCase {
return &GetTodoUseCase{
repository: repo,
}
}

func (uc *GetTodoUseCase) Execute(ctx context.Context, input dto.GetTodoInput) (*dto.GetTodoOutput, error) {
todo, err := uc.repository.GetByID(ctx, input.ID)
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
