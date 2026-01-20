package usecases

import (
	"context"
	"test-go/internal/features/todo/dto"
	"test-go/internal/features/todo/usecase"
	"test-go/internal/features/todo/entity"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestUpdateTodoUseCase_Execute(t *testing.T) {
	repo := NewMockTodoRepository()

	testID := primitive.NewObjectID()
	todo := &entity.Todo{
		ID:          testID,
		Title:       "Original title",
		Description: "Original description",
		Completed:   false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	repo.Create(context.Background(), todo)

	useCase := usecase.NewUpdateTodoUseCase(repo)
	input := dto.UpdateTodoInput{
		ID:          testID.Hex(),
		Title:       "Updated title",
		Description: "Updated description",
		Completed:   true,
	}

	output, err := useCase.Execute(context.Background(), input)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if output.Title != input.Title {
		t.Errorf("Expected title %s, got %s", input.Title, output.Title)
	}

	if output.Description != input.Description {
		t.Errorf("Expected description %s, got %s", input.Description, output.Description)
	}

	if output.Completed != input.Completed {
		t.Errorf("Expected completed %v, got %v", input.Completed, output.Completed)
	}

	if output.ID != testID {
		t.Errorf("Expected ID %s, got %s", testID.Hex(), output.ID.Hex())
	}
}

func TestUpdateTodoUseCase_NotFound(t *testing.T) {
	repo := NewMockTodoRepository()
	useCase := usecase.NewUpdateTodoUseCase(repo)
	input := dto.UpdateTodoInput{
		ID:    "nonexistent",
		Title: "Updated title",
	}

	output, err := useCase.Execute(context.Background(), input)

	if err == nil {
		t.Fatal("Expected error for nonexistent todo, got nil")
	}

	if output != nil {
		t.Error("Expected output to be nil")
	}
}
