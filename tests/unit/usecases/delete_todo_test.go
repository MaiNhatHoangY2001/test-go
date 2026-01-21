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

func TestDeleteTodoUseCase_Execute(t *testing.T) {
	repo := NewMockTodoRepository()

	userID := "test-user-id"
	testID := primitive.NewObjectID()
	todo := &entities.Todo{
		ID:          testID,
		UserID:      userID,
		Title:       "Test todo",
		Description: "Test",
		Completed:   false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	repo.Create(context.Background(), todo)

	useCase := usecase.NewDeleteTodoUseCase(repo)
	input := dto.DeleteTodoInput{ID: testID.Hex()}

	err := useCase.Execute(context.Background(), userID, input)

	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Verify the todo is deleted
	_, err = repo.GetByID(context.Background(), userID, testID.Hex())
	if err == nil {
		t.Error("Expected error when getting deleted todo, but got nil")
	}
}

func TestDeleteTodoUseCase_NotFound(t *testing.T) {
	repo := NewMockTodoRepository()
	useCase := usecase.NewDeleteTodoUseCase(repo)
	userID := "test-user-id"
	input := dto.DeleteTodoInput{ID: "nonexistent"}

	err := useCase.Execute(context.Background(), userID, input)

	if err == nil {
		t.Fatal("Expected error for nonexistent todo, got nil")
	}
}
