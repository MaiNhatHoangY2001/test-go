package handler

import (
	"test-go/internal/features/todo/dto"
	"test-go/internal/features/todo/usecase"
	"test-go/internal/shared/response"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type TodoHandler struct {
	createUseCase  *usecase.CreateTodoUseCase
	getTodoUseCase *usecase.GetTodoUseCase
	getAllUseCase  *usecase.GetAllTodosUseCase
	updateUseCase  *usecase.UpdateTodoUseCase
	deleteUseCase  *usecase.DeleteTodoUseCase
	logger         *logrus.Logger
}

func NewTodoHandler(
	createUC *usecase.CreateTodoUseCase,
	getTodoUC *usecase.GetTodoUseCase,
	getAllUC *usecase.GetAllTodosUseCase,
	updateUC *usecase.UpdateTodoUseCase,
	deleteUC *usecase.DeleteTodoUseCase,
	logger *logrus.Logger,
) *TodoHandler {
	return &TodoHandler{
		createUseCase:  createUC,
		getTodoUseCase: getTodoUC,
		getAllUseCase:  getAllUC,
		updateUseCase:  updateUC,
		deleteUseCase:  deleteUC,
		logger:         logger,
	}
}

// CreateTodo godoc
//
//	@Summary		Create a new todo
//	@Description	Create a new todo for the authenticated user
//	@Tags			todos
//	@Accept			json
//	@Produce		json
//	@Param			todo	body		dto.CreateTodoInput								true	"Todo object"
//	@Success		201		{object}	response.APIResponse{data=dto.CreateTodoOutput}	"Successfully created"
//	@Failure		400		{object}	response.APIResponse{error=response.ErrorInfo}	"Bad request"
//	@Failure		401		{object}	response.APIResponse{error=response.ErrorInfo}	"Unauthorized"
//	@Failure		500		{object}	response.APIResponse{error=response.ErrorInfo}	"Internal server error"
//	@Security		BearerAuth
//	@Router			/todos [post]
func (h *TodoHandler) CreateTodo(c *gin.Context) {
	// Get user ID from context (set by auth middleware)
	userID, exists := c.Get("user_id")
	if !exists {
		h.logger.WithFields(logrus.Fields{"request_id": c.GetString("X-Request-ID")}).Error("user_id_not_found_in_context")
		response.Unauthorized(c, "Unauthorized")
		return
	}

	var input dto.CreateTodoInput

	if err := c.ShouldBindJSON(&input); err != nil {
		h.logger.WithFields(logrus.Fields{"request_id": c.GetString("X-Request-ID")}).Warn("invalid_request: " + err.Error())
		response.BadRequest(c, "Invalid request body")
		return
	}

	output, err := h.createUseCase.Execute(c.Request.Context(), userID.(string), input)
	if err != nil {
		h.logger.WithFields(logrus.Fields{"request_id": c.GetString("X-Request-ID")}).Error("create_todo_failed: " + err.Error())
		response.HandleError(c, h.logger, err)
		return
	}

	h.logger.WithFields(logrus.Fields{"request_id": c.GetString("X-Request-ID"), "todo_id": output.ID}).Info("todo_created")
	response.Created(c, output)
}

// GetAllTodos godoc
//
//	@Summary		Get all todos
//	@Description	Get all todos for the authenticated user with pagination
//	@Tags			todos
//	@Accept			json
//	@Produce		json
//	@Param			page	query		int																				false	"Page number"	default(1)
//	@Param			limit	query		int																				false	"Page size"		default(10)
//	@Success		200		{object}	response.PagingResponse{data=response.PagingData{data=[]dto.GetAllTodosOutput}}	"Successfully retrieved"
//	@Security		BearerAuth
//	@Router			/todos [get]
func (h *TodoHandler) GetAllTodos(c *gin.Context) {
	// Get user ID from context (set by auth middleware)
	userID, exists := c.Get("user_id")
	if !exists {
		h.logger.WithFields(logrus.Fields{"request_id": c.GetString("X-Request-ID")}).Error("user_id_not_found_in_context")
		response.Unauthorized(c, "Unauthorized")
		return
	}

	var input dto.GetAllTodosInput
	if err := c.ShouldBindQuery(&input); err != nil {
		h.logger.WithFields(logrus.Fields{"request_id": c.GetString("X-Request-ID")}).Warn("invalid_query_params: " + err.Error())
		response.BadRequest(c, "Invalid query parameters")
		return
	}

	result, err := h.getAllUseCase.Execute(c.Request.Context(), userID.(string), input)
	if err != nil {
		h.logger.WithFields(logrus.Fields{"request_id": c.GetString("X-Request-ID")}).Error("get_todos_failed: " + err.Error())
		response.HandleError(c, h.logger, err)
		return
	}

	h.logger.WithFields(logrus.Fields{
		"request_id": c.GetString("X-Request-ID"),
		"count":      len(result.Data),
		"page":       result.Pagination.PageNum,
		"total":      result.Pagination.TotalItems,
	}).Info("todos_retrieved")

	// Use paging response format
	response.SuccessPaging(c,
		int64(result.Pagination.PageNum),
		int64(result.Pagination.PageSize),
		result.Pagination.TotalItems,
		result.Data,
	)
}

func (h *TodoHandler) GetTodo(c *gin.Context) {
	// Get user ID from context (set by auth middleware)
	userID, exists := c.Get("user_id")
	if !exists {
		h.logger.WithFields(logrus.Fields{"request_id": c.GetString("X-Request-ID")}).Error("user_id_not_found_in_context")
		response.Unauthorized(c, "Unauthorized")
		return
	}

	id := c.Param("id")

	output, err := h.getTodoUseCase.Execute(c.Request.Context(), userID.(string), dto.GetTodoInput{ID: id})
	if err != nil {
		h.logger.WithFields(logrus.Fields{"request_id": c.GetString("X-Request-ID"), "todo_id": id}).Warn("todo_not_found: " + err.Error())
		response.HandleError(c, h.logger, err)
		return
	}

	h.logger.WithFields(logrus.Fields{"request_id": c.GetString("X-Request-ID"), "todo_id": id}).Info("todo_retrieved")
	response.OK(c, output)
}

func (h *TodoHandler) UpdateTodo(c *gin.Context) {
	// Get user ID from context (set by auth middleware)
	userID, exists := c.Get("user_id")
	if !exists {
		h.logger.WithFields(logrus.Fields{"request_id": c.GetString("X-Request-ID")}).Error("user_id_not_found_in_context")
		response.Unauthorized(c, "Unauthorized")
		return
	}

	id := c.Param("id")
	var input dto.UpdateTodoInput

	if err := c.ShouldBindJSON(&input); err != nil {
		h.logger.WithFields(logrus.Fields{"request_id": c.GetString("X-Request-ID"), "todo_id": id}).Warn("invalid_request: " + err.Error())
		response.BadRequest(c, "Invalid request body")
		return
	}

	input.ID = id
	output, err := h.updateUseCase.Execute(c.Request.Context(), userID.(string), input)
	if err != nil {
		h.logger.WithFields(logrus.Fields{"request_id": c.GetString("X-Request-ID"), "todo_id": id}).Error("update_todo_failed: " + err.Error())
		response.HandleError(c, h.logger, err)
		return
	}

	h.logger.WithFields(logrus.Fields{"request_id": c.GetString("X-Request-ID"), "todo_id": id}).Info("todo_updated")
	response.OK(c, output)
}

func (h *TodoHandler) DeleteTodo(c *gin.Context) {
	// Get user ID from context (set by auth middleware)
	userID, exists := c.Get("user_id")
	if !exists {
		h.logger.WithFields(logrus.Fields{"request_id": c.GetString("X-Request-ID")}).Error("user_id_not_found_in_context")
		response.Unauthorized(c, "Unauthorized")
		return
	}

	id := c.Param("id")
	err := h.deleteUseCase.Execute(c.Request.Context(), userID.(string), dto.DeleteTodoInput{ID: id})
	if err != nil {
		h.logger.WithFields(logrus.Fields{"request_id": c.GetString("X-Request-ID"), "todo_id": id}).Error("delete_todo_failed: " + err.Error())
		response.HandleError(c, h.logger, err)
		return
	}

	h.logger.WithFields(logrus.Fields{"request_id": c.GetString("X-Request-ID"), "todo_id": id}).Info("todo_deleted")
	response.NoContent(c)
}
