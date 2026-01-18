package controllers

import (
	"net/http"
	"strconv"
	"todo-app/internal/models"
	"todo-app/pkg/database"

	"github.com/gin-gonic/gin"
)

type TodoController struct{}

func NewTodoController() *TodoController {
	return &TodoController{}
}

func (tc *TodoController) GetAllTodos(c *gin.Context) {
	todos := database.GetAllTodos()

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   todos,
		"count":  len(todos),
	})
}

func (tc *TodoController) CreateTodo(c *gin.Context) {
	var todoReq models.TodoRequest

	if err := c.ShouldBindJSON(&todoReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Invalid request data",
			"error":   err.Error(),
		})

		return
	}

	todo := database.AddTodo(todoReq)

	c.JSON(http.StatusCreated, gin.H{
		"status":  "success",
		"message": "Todo created successfully",
		"data":    todo,
	})
}

func (tc *TodoController) GetTodoByID(c *gin.Context) {
	idStr, exist := c.Params.Get("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil || !exist {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Invalid todo ID",
		})
		return
	}

	todo, found := database.GetTodoByID(uint(id))
	if !found {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "error",
			"message": "Todo not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": "success",
		"data":   todo,
	})
}

func (tc *TodoController) UpdateTodo(c *gin.Context) {
	idStr, exist := c.Params.Get("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil || !exist {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Invalid todo ID",
		})
		return
	}

	var todoReq models.TodoRequest
	if err := c.ShouldBindJSON(&todoReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Invalid request data",
			"error":   err.Error(),
		})
		return
	}

	todo, found := database.UpdateTodo(uint(id), todoReq)
	if !found {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "error",
			"message": "Todo not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Todo updated successfully",
		"data":    todo,
	})
}

func (tc *TodoController) DeleteTodo(c *gin.Context) {
	idStr, exist := c.Params.Get("id")
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil || !exist {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  "error",
			"message": "Invalid todo ID",
		})
		return
	}

	if !database.DeleteTodo(uint(id)) {
		c.JSON(http.StatusNotFound, gin.H{
			"status":  "error",
			"message": "Todo not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  "success",
		"message": "Todo deleted successfully",
	})
}
