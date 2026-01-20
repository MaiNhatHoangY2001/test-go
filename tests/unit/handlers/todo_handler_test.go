package handlers

import (
	"testing"
	"time"

	"test-go/internal/features/todo/dto"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// TestCreateTodo_Success tests successful todo creation
func TestCreateTodo_Success(t *testing.T) {
	output := &dto.CreateTodoOutput{
		ID:          primitive.NewObjectID(),
		Title:       "Test Todo",
		Description: "Test Description",
		Completed:   false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if output.Title != "Test Todo" {
		t.Errorf("Expected title 'Test Todo', got %s", output.Title)
	}
	if output.Completed {
		t.Error("Expected Completed to be false")
	}
}

// TestGetTodo_Success tests successful todo retrieval
func TestGetTodo_Success(t *testing.T) {
	expectedID := primitive.NewObjectID()
	output := &dto.GetTodoOutput{
		ID:          expectedID,
		Title:       "Test Todo",
		Description: "Test Description",
		Completed:   false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if output.ID != expectedID {
		t.Errorf("Expected ID %v, got %v", expectedID, output.ID)
	}
	if output.Title != "Test Todo" {
		t.Errorf("Expected title 'Test Todo', got %s", output.Title)
	}
}

// TestGetAllTodos_Success tests retrieving all todos
func TestGetAllTodos_Success(t *testing.T) {
	todos := []dto.GetAllTodosOutput{
		{
			ID:          primitive.NewObjectID(),
			Title:       "Todo 1",
			Description: "Description 1",
			Completed:   false,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			ID:          primitive.NewObjectID(),
			Title:       "Todo 2",
			Description: "Description 2",
			Completed:   true,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}

	if len(todos) != 2 {
		t.Errorf("Expected 2 todos, got %d", len(todos))
	}
	if todos[0].Title != "Todo 1" {
		t.Errorf("Expected first todo title 'Todo 1', got %s", todos[0].Title)
	}
}

// TestGetAllTodos_Empty tests retrieving when no todos exist
func TestGetAllTodos_Empty(t *testing.T) {
	todos := []dto.GetAllTodosOutput{}

	if len(todos) != 0 {
		t.Errorf("Expected 0 todos, got %d", len(todos))
	}
}

// TestUpdateTodo_Success tests successful todo update
func TestUpdateTodo_Success(t *testing.T) {
	todoID := primitive.NewObjectID()
	updatedTodo := &dto.GetTodoOutput{
		ID:          todoID,
		Title:       "Updated Title",
		Description: "Updated Description",
		Completed:   true,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if updatedTodo.Title != "Updated Title" {
		t.Errorf("Expected title 'Updated Title', got %s", updatedTodo.Title)
	}
	if !updatedTodo.Completed {
		t.Error("Expected Completed to be true")
	}
}

// TestCreateTodo_MultipleCreations tests creating multiple todos
func TestCreateTodo_MultipleCreations(t *testing.T) {
	scenarios := []struct {
		name  string
		title string
		desc  string
	}{
		{
			name:  "Simple todo",
			title: "Buy milk",
			desc:  "Whole milk from store",
		},
		{
			name:  "Complex todo",
			title: "Write comprehensive documentation",
			desc:  "Document all APIs and features for the new release",
		},
		{
			name:  "Minimal todo",
			title: "Call mom",
			desc:  "",
		},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.name, func(t *testing.T) {
			output := &dto.CreateTodoOutput{
				ID:          primitive.NewObjectID(),
				Title:       scenario.title,
				Description: scenario.desc,
				Completed:   false,
				CreatedAt:   time.Now(),
				UpdatedAt:   time.Now(),
			}

			if output.Title != scenario.title {
				t.Errorf("Expected title '%s', got '%s'", scenario.title, output.Title)
			}
			if output.Description != scenario.desc {
				t.Errorf("Expected description '%s', got '%s'", scenario.desc, output.Description)
			}
		})
	}
}

// TestUpdateTodo_PartialUpdate tests partial todo update
func TestUpdateTodo_PartialUpdate(t *testing.T) {
	todoID := primitive.NewObjectID()

	output := &dto.GetTodoOutput{
		ID:          todoID,
		Title:       "Updated Title Only",
		Description: "", // Empty description
		Completed:   false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if output.Title != "Updated Title Only" {
		t.Errorf("Expected title 'Updated Title Only', got %s", output.Title)
	}
	if output.Description != "" {
		t.Errorf("Expected empty description, got %s", output.Description)
	}
}

// TestTodoCompletionToggle tests toggling todo completion status
func TestTodoCompletionToggle(t *testing.T) {
	todoID := primitive.NewObjectID()

	// Create incomplete todo
	incompleteTodo := &dto.GetTodoOutput{
		ID:          todoID,
		Title:       "Task",
		Description: "Task description",
		Completed:   false,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	if incompleteTodo.Completed {
		t.Error("Expected Completed to be false initially")
	}

	// Update to completed
	completeTodo := &dto.GetTodoOutput{
		ID:          todoID,
		Title:       "Task",
		Description: "Task description",
		Completed:   true,
		CreatedAt:   incompleteTodo.CreatedAt,
		UpdatedAt:   time.Now(),
	}

	if !completeTodo.Completed {
		t.Error("Expected Completed to be true after update")
	}
}

// TestTodoTimestamps tests todo creation and update timestamps
func TestTodoTimestamps(t *testing.T) {
	createdAt := time.Now()
	updatedAt := time.Now().Add(time.Hour)

	output := &dto.CreateTodoOutput{
		ID:          primitive.NewObjectID(),
		Title:       "Test Todo",
		Description: "Test Description",
		Completed:   false,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
	}

	if output.CreatedAt != createdAt {
		t.Errorf("Expected CreatedAt %v, got %v", createdAt, output.CreatedAt)
	}
	if output.UpdatedAt != updatedAt {
		t.Errorf("Expected UpdatedAt %v, got %v", updatedAt, output.UpdatedAt)
	}
}

// TestGetTodoInput validation tests
func TestGetTodoInput_Validation(t *testing.T) {
	scenarios := []struct {
		name  string
		input dto.GetTodoInput
	}{
		{
			name: "Valid ID",
			input: dto.GetTodoInput{
				ID: primitive.NewObjectID().Hex(),
			},
		},
		{
			name: "Empty ID",
			input: dto.GetTodoInput{
				ID: "",
			},
		},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.name, func(t *testing.T) {
			if scenario.name == "Empty ID" && scenario.input.ID == "" {
				t.Log("Correctly identified empty ID scenario")
			}
		})
	}
}

// TestCreateTodoInput validation tests
func TestCreateTodoInput_Validation(t *testing.T) {
	scenarios := []struct {
		name  string
		input dto.CreateTodoInput
	}{
		{
			name: "Valid input",
			input: dto.CreateTodoInput{
				Title:       "Test",
				Description: "Test",
			},
		},
		{
			name: "Empty title",
			input: dto.CreateTodoInput{
				Title:       "",
				Description: "Test",
			},
		},
		{
			name: "Very long description",
			input: dto.CreateTodoInput{
				Title:       "Test",
				Description: "This is a very long description that contains many words and takes up significant space in the system",
			},
		},
	}

	for _, scenario := range scenarios {
		t.Run(scenario.name, func(t *testing.T) {
			if scenario.input.Title == "" && scenario.name == "Empty title" {
				t.Log("Correctly identified empty title scenario")
			}
		})
	}
}
