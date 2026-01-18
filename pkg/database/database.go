package database

import (
	"sync"
	"time"
	"todo-app/internal/models"
)

var (
	todos   []models.Todo
	nextID  uint = 1
	todoMux sync.RWMutex
)

func InitDB() {
	todos = make([]models.Todo, 0)
}

func GetAllTodos() []models.Todo {
	todoMux.RLock()
	defer todoMux.RUnlock()
	return todos
}

func AddTodo(todoReq models.TodoRequest) models.Todo {
	todoMux.Lock()
	defer todoMux.Unlock()

	todo := models.Todo{
		ID:          nextID,
		Title:       todoReq.Title,
		Description: todoReq.Description,
		Completed:   todoReq.Completed,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	todos = append(todos, todo)
	nextID++

	return todo
}

func GetTodoByID(id uint) (*models.Todo, bool) {
	todoMux.RLock()
	defer todoMux.RUnlock()

	for i := range todos {
		if todos[i].ID == id {
			return &todos[i], true
		}
	}

	return nil, false
}

func UpdateTodo(id uint, todoReq models.TodoRequest) (*models.Todo, bool) {
	todoMux.Lock()
	defer todoMux.Unlock()

	for i := range todos {
		if todos[i].ID == id {
			todos[i].Title = todoReq.Title
			todos[i].Description = todoReq.Description
			todos[i].Completed = todoReq.Completed
			todos[i].UpdatedAt = time.Now()
			return &todos[i], true
		}
	}

	return nil, false
}

func DeleteTodo(id uint) bool {
	todoMux.Lock()
	defer todoMux.Unlock()

	for i := range todos {
		if todos[i].ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			return true
		}
	}
	return false
}
