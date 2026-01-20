package usecases

import (
	"context"
	"test-go/internal/application/usecases"
	"test-go/internal/domain/entities"
	"test-go/internal/infrastructure/repository"
	"testing"
	"time"
)

func TestGetTodoUseCase_Execute(t *testing.T) {
	repo := repository.NewInMemoryTodoRepository()

	todo := &entities.Todo{
		ID:          "test-id",
		Title:       "Test todo",
		Description: "Test",
		Completed:   false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	repo.Create(context.Background(), todo)

	useCase := usecases.NewGetTodoUseCase(repo)
	input := usecases.GetTodoInput{ID: "test-id"}

	output, err := useCase.Execute(context.Background(), input)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if output.ID != "test-id" {
		t.Errorf("Expected ID test-id, got %s", output.ID)
	}

	if output.Title != "Test todo" {
		t.Errorf("Expected title Test Todo, got %s", output.Title)
	}
}

func TestGetTodoUseCase_NotFound(t *testing.T) {
	repo := repository.NewInMemoryTodoRepository()
	useCase := usecases.NewGetTodoUseCase(repo)
	input := usecases.GetTodoInput{ID: "nonexistent"}
	output, err := useCase.Execute(context.Background(), input)

	if err == nil {
		t.Fatal("Expected error, got nil")
	}

	if output != nil {
		t.Error("Expected output to be nil")
	}
}
