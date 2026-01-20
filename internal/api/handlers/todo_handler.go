package handlers

import (
	"net/http"
	"test-go/internal/application/usecases"

	"github.com/gin-gonic/gin"
)

type TodoHandler struct {
	createUseCase  *usecases.CreateTodoUseCase
	getTodoUseCase *usecases.GetTodoUseCase
	getAllUseCase  *usecases.GetAllTodosUseCase
	updateUseCase  *usecases.UpdateTodoUseCase
	deleteUseCase  *usecases.DeleteTodoUseCase
}

func NewTodoHandler(
	createUC *usecases.CreateTodoUseCase,
	getTodoUC *usecases.GetTodoUseCase,
	getAllUC *usecases.GetAllTodosUseCase,
	updateUC *usecases.UpdateTodoUseCase,
	deleteUC *usecases.DeleteTodoUseCase,
) *TodoHandler {
	return &TodoHandler{
		createUseCase:  createUC,
		getTodoUseCase: getTodoUC,
		getAllUseCase:  getAllUC,
		updateUseCase:  updateUC,
		deleteUseCase:  deleteUC,
	}
}

func (h *TodoHandler) CreateTodo(c *gin.Context) {
	var input usecases.CreateTodoInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	output, err := h.createUseCase.Execute(c.Request.Context(), input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, output)
}

func (h *TodoHandler) GetAllTodos(c *gin.Context) {
	outputs, err := h.getAllUseCase.Execute(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, outputs)
}

func (h *TodoHandler) GetTodo(c *gin.Context) {
	id, _ := c.Params.Get("id")

	output, err := h.getTodoUseCase.Execute(c.Request.Context(), usecases.GetTodoInput{ID: id})
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, output)
}

func (h *TodoHandler) UpdateTodo(c *gin.Context) {
	id, _ := c.Params.Get("id")
	var input usecases.UpdateTodoInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input.ID = id
	output, err := h.updateUseCase.Execute(c.Request.Context(), input)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, output)
}

func (h *TodoHandler) DeleteTodo(c *gin.Context) {
	id, _ := c.Params.Get("id")

	err := h.deleteUseCase.Execute(c.Request.Context(), usecases.DeleteTodoInput{ID: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
