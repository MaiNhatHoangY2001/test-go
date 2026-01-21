package dto

import (
	"time"

	sharedDto "test-go/internal/shared/dto"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CreateTodoInput struct {
	Title       string `json:"title" binding:"required,min=1,max=200"`
	Description string `json:"description" binding:"max=1000"`
}

type CreateTodoOutput struct {
	ID          primitive.ObjectID `json:"id"`
	Title       string             `json:"title"`
	Description string             `json:"description"`
	Completed   bool               `json:"completed"`
	CreatedAt   time.Time          `json:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at"`
}

type GetTodoInput struct {
	ID string
}

type GetTodoOutput struct {
	ID          primitive.ObjectID `json:"id"`
	Title       string             `json:"title"`
	Description string             `json:"description"`
	Completed   bool               `json:"completed"`
	CreatedAt   time.Time          `json:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at"`
}

type GetAllTodosInput struct {
	sharedDto.PaginationInput
}

type GetAllTodosOutput struct {
	ID          primitive.ObjectID `json:"id"`
	Title       string             `json:"title"`
	Description string             `json:"description"`
	Completed   bool               `json:"completed"`
	CreatedAt   time.Time          `json:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at"`
}

type GetAllTodosResponse struct {
	Data       []GetAllTodosOutput      `json:"data"`
	Pagination sharedDto.PaginationInfo `json:"pagination"`
}

type UpdateTodoInput struct {
	ID          string `json:"-"`
	Title       string `json:"title" binding:"required,min=1,max=200"`
	Description string `json:"description" binding:"max=1000"`
	Completed   bool   `json:"completed"`
}

type UpdateTodoOutput struct {
	ID          primitive.ObjectID `json:"id"`
	Title       string             `json:"title"`
	Description string             `json:"description"`
	Completed   bool               `json:"completed"`
	CreatedAt   time.Time          `json:"created_at"`
	UpdatedAt   time.Time          `json:"updated_at"`
}

type DeleteTodoInput struct {
	ID string
}
