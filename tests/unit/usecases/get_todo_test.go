package usecases

import (
	"context"
	"test-go/internal/domain/entities"
	"test-go/internal/features/todo/dto"
	"test-go/internal/features/todo/usecase"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestGetTodoUseCase_Execute(t *testing.T) {
	repo := NewMockTodoRepository()

	testID := primitive.NewObjectID()
	todo := &entities.Todo{
		ID:          testID,
		Title:       "Test todo",
		Description: "Test",
		Completed:   false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	repo.Create(context.Background(), todo)

	useCase := usecase.NewGetTodoUseCase(repo)
	input := dto.GetTodoInput{ID: testID.Hex()}

	output, err := useCase.Execute(context.Background(), input)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if output.ID != testID {
		t.Errorf("Expected ID %s, got %s", testID.Hex(), output.ID.Hex())
	}

	if output.Title != "Test todo" {
		t.Errorf("Expected title Test todo, got %s", output.Title)
	}
}

func TestGetTodoUseCase_NotFound(t *testing.T) {
	repo := NewMockTodoRepository()
	useCase := usecase.NewGetTodoUseCase(repo)
	input := dto.GetTodoInput{ID: "nonexistent"}
	output, err := useCase.Execute(context.Background(), input)

	if err == nil {
		t.Fatal("Expected error, got nil")
	}

	if output != nil {
		t.Error("Expected output to be nil")
	}
}
