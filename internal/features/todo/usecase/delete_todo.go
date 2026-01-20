package usecase

import (
"context"
"test-go/internal/features/todo/dto"
"test-go/internal/features/todo/repository"
)

type DeleteTodoUseCase struct {
repository repository.TodoRepository
}

func NewDeleteTodoUseCase(repo repository.TodoRepository) *DeleteTodoUseCase {
return &DeleteTodoUseCase{
repository: repo,
}
}

func (uc *DeleteTodoUseCase) Execute(ctx context.Context, input dto.DeleteTodoInput) error {
return uc.repository.Delete(ctx, input.ID)
}
