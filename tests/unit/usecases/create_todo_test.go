package usecases

import (
	"context"
	"test-go/internal/application/usecases"
	"test-go/internal/infrastructure/repository"
	"testing"
)

func TestCreateTodoUseCase_Execute(t *testing.T) {

	repo := repository.NewInMemoryTodoRepository()
	useCase := usecases.NewCreateTodoUseCase(repo)

	input := usecases.CreateTodoInput{
		Title:       "Test Todo",
		Description: "Test Description",
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

	if output.ID == "" {
		t.Error("Expected ID to be set")
	}

	if output.Completed != false {
		t.Error("Expected completed to be false")
	}
}
