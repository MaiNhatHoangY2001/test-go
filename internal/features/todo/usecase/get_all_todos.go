package usecase

import (
"context"
"test-go/internal/features/todo/dto"
"test-go/internal/features/todo/repository"
)

type GetAllTodosUseCase struct {
repository repository.TodoRepository
}

func NewGetAllTodosUseCase(repo repository.TodoRepository) *GetAllTodosUseCase {
return &GetAllTodosUseCase{
repository: repo,
}
}

func (uc *GetAllTodosUseCase) Execute(ctx context.Context) ([]dto.GetAllTodosOutput, error) {
todos, err := uc.repository.GetAll(ctx)
if err != nil {
return nil, err
}

outputs := make([]dto.GetAllTodosOutput, len(todos))
for i, todo := range todos {
outputs[i] = dto.GetAllTodosOutput{
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
