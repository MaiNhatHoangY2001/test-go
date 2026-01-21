package usecase

import (
	"context"
	"test-go/internal/domain/repositories"
	"test-go/internal/features/todo/dto"
	sharedDto "test-go/internal/shared/dto"
	"test-go/pkg/constants"
)

type GetAllTodosUseCase struct {
	repository repositories.TodoRepository
}

func NewGetAllTodosUseCase(repo repositories.TodoRepository) *GetAllTodosUseCase {
return &GetAllTodosUseCase{
repository: repo,
}
}

func (uc *GetAllTodosUseCase) Execute(ctx context.Context, userID string, input dto.GetAllTodosInput) (*dto.GetAllTodosResponse, error) {
	// Set default pagination values
	pageNum := input.PageNum
	if pageNum <= 0 {
		pageNum = constants.DefaultPageNum
	}
	pageSize := input.PageSize
	if pageSize <= 0 {
		pageSize = constants.DefaultPageSize
	}

	todos, totalCount, err := uc.repository.GetAll(ctx, userID, pageNum, pageSize)
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

	totalPages := int(totalCount) / pageSize
	if int(totalCount)%pageSize != 0 {
		totalPages++
	}

	return &dto.GetAllTodosResponse{
		Data: outputs,
		Pagination: sharedDto.PaginationInfo{
			PageNum:    pageNum,
			PageSize:   pageSize,
			TotalItems: totalCount,
			TotalPages: totalPages,
		},
	}, nil
}
