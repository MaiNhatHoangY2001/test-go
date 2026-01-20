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

func (h *TodoHandler) CreateTodo(c *gin.Context) {
var input dto.CreateTodoInput

if err := c.ShouldBindJSON(&input); err != nil {
h.logger.WithFields(logrus.Fields{"request_id": c.GetString("X-Request-ID")}).Warn("invalid_request: " + err.Error())
response.BadRequest(c, "Invalid request body")
return
}

output, err := h.createUseCase.Execute(c.Request.Context(), input)
if err != nil {
h.logger.WithFields(logrus.Fields{"request_id": c.GetString("X-Request-ID")}).Error("create_todo_failed: " + err.Error())
response.HandleError(c, h.logger, err)
return
}

h.logger.WithFields(logrus.Fields{"request_id": c.GetString("X-Request-ID"), "todo_id": output.ID}).Info("todo_created")
response.Created(c, output)
}

func (h *TodoHandler) GetAllTodos(c *gin.Context) {
outputs, err := h.getAllUseCase.Execute(c.Request.Context())
if err != nil {
h.logger.WithFields(logrus.Fields{"request_id": c.GetString("X-Request-ID")}).Error("get_todos_failed: " + err.Error())
response.HandleError(c, h.logger, err)
return
}

h.logger.WithFields(logrus.Fields{"request_id": c.GetString("X-Request-ID"), "count": len(outputs)}).Info("todos_retrieved")
response.OK(c, outputs)
}

func (h *TodoHandler) GetTodo(c *gin.Context) {
id := c.Param("id")

output, err := h.getTodoUseCase.Execute(c.Request.Context(), dto.GetTodoInput{ID: id})
if err != nil {
h.logger.WithFields(logrus.Fields{"request_id": c.GetString("X-Request-ID"), "todo_id": id}).Warn("todo_not_found: " + err.Error())
response.HandleError(c, h.logger, err)
return
}

h.logger.WithFields(logrus.Fields{"request_id": c.GetString("X-Request-ID"), "todo_id": id}).Info("todo_retrieved")
response.OK(c, output)
}

func (h *TodoHandler) UpdateTodo(c *gin.Context) {
id := c.Param("id")
var input dto.UpdateTodoInput

if err := c.ShouldBindJSON(&input); err != nil {
h.logger.WithFields(logrus.Fields{"request_id": c.GetString("X-Request-ID"), "todo_id": id}).Warn("invalid_request: " + err.Error())
response.BadRequest(c, "Invalid request body")
return
}

input.ID = id
output, err := h.updateUseCase.Execute(c.Request.Context(), input)
if err != nil {
h.logger.WithFields(logrus.Fields{"request_id": c.GetString("X-Request-ID"), "todo_id": id}).Error("update_todo_failed: " + err.Error())
response.HandleError(c, h.logger, err)
return
}

h.logger.WithFields(logrus.Fields{"request_id": c.GetString("X-Request-ID"), "todo_id": id}).Info("todo_updated")
response.OK(c, output)
}

func (h *TodoHandler) DeleteTodo(c *gin.Context) {
id := c.Param("id")
err := h.deleteUseCase.Execute(c.Request.Context(), dto.DeleteTodoInput{ID: id})
if err != nil {
h.logger.WithFields(logrus.Fields{"request_id": c.GetString("X-Request-ID"), "todo_id": id}).Error("delete_todo_failed: " + err.Error())
response.HandleError(c, h.logger, err)
return
}

h.logger.WithFields(logrus.Fields{"request_id": c.GetString("X-Request-ID"), "todo_id": id}).Info("todo_deleted")
response.NoContent(c)
}
