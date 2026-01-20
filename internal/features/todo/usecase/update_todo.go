package usecase

import (
"context"
"test-go/internal/features/todo/dto"
"test-go/internal/features/todo/repository"
"time"
)

type UpdateTodoUseCase struct {
repository repository.TodoRepository
}

func NewUpdateTodoUseCase(repo repository.TodoRepository) *UpdateTodoUseCase {
return &UpdateTodoUseCase{
repository: repo,
}
}

func (uc *UpdateTodoUseCase) Execute(ctx context.Context, input dto.UpdateTodoInput) (*dto.UpdateTodoOutput, error) {
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

return &dto.UpdateTodoOutput{
ID:          existingTodo.ID,
Title:       existingTodo.Title,
Description: existingTodo.Description,
Completed:   existingTodo.Completed,
CreatedAt:   existingTodo.CreatedAt,
UpdatedAt:   existingTodo.UpdatedAt,
}, nil
}
