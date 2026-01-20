package usecase

import (
	"context"
	"test-go/internal/domain/repositories"
	"test-go/internal/features/todo/dto"
)

type GetAllTodosUseCase struct {
	repository repositories.TodoRepository
}

func NewGetAllTodosUseCase(repo repositories.TodoRepository) *GetAllTodosUseCase {
return &GetAllTodosUseCase{
repository: repo,
}
}

func (uc *GetAllTodosUseCase) Execute(ctx context.Context, input dto.GetAllTodosInput) (*dto.GetAllTodosResponse, error) {
	// Set default pagination values
	page := input.Page
	if page <= 0 {
		page = 1
	}
	limit := input.Limit
	if limit <= 0 {
		limit = 10
	}

	todos, totalCount, err := uc.repository.GetAll(ctx, page, limit)
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

	totalPages := int(totalCount) / limit
	if int(totalCount)%limit != 0 {
		totalPages++
	}

	return &dto.GetAllTodosResponse{
		Data: outputs,
		Pagination: dto.PaginationInfo{
			Page:       page,
			Limit:      limit,
			TotalItems: totalCount,
			TotalPages: totalPages,
		},
	}, nil
}
