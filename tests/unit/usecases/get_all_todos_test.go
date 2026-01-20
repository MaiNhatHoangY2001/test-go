package usecases

import (
	"context"
	"test-go/internal/domain/entities"
	"test-go/internal/features/todo/usecase"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func TestGetAllTodosUseCase_Execute(t *testing.T) {
	repo := NewMockTodoRepository()

	// Create test todos
	todo1 := &entities.Todo{
		ID:          primitive.NewObjectID(),
		Title:       "Todo 1",
		Description: "Description 1",
		Completed:   false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	todo2 := &entities.Todo{
		ID:          primitive.NewObjectID(),
		Title:       "Todo 2",
		Description: "Description 2",
		Completed:   true,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	repo.Create(context.Background(), todo1)
	repo.Create(context.Background(), todo2)

	useCase := usecase.NewGetAllTodosUseCase(repo)
	outputs, err := useCase.Execute(context.Background())

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(outputs) != 2 {
		t.Errorf("Expected 2 todos, got %d", len(outputs))
	}

	// Verify both todos are present (order may vary due to map iteration)
	titles := make(map[string]bool)
	for _, output := range outputs {
		titles[output.Title] = true
	}

	if !titles["Todo 1"] {
		t.Error("Expected Todo 1 to be in results")
	}

	if !titles["Todo 2"] {
		t.Error("Expected Todo 2 to be in results")
	}
}

func TestGetAllTodosUseCase_EmptyList(t *testing.T) {
	repo := NewMockTodoRepository()
	useCase := usecase.NewGetAllTodosUseCase(repo)

	outputs, err := useCase.Execute(context.Background())

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	if len(outputs) != 0 {
		t.Errorf("Expected 0 todos, got %d", len(outputs))
	}
}
