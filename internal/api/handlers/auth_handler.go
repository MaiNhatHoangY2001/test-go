package handlers

import (
	"net/http"
	"test-go/internal/application/dto"
	"test-go/internal/application/usecases"

	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	signUpUseCase *usecases.SignUpUseCase
	loginUseCase  *usecases.LoginUseCase
}

func NewAuthHandler(
	signUpUseCase *usecases.SignUpUseCase,
	loginUseCase *usecases.LoginUseCase,
) *AuthHandler {
	return &AuthHandler{
		signUpUseCase: signUpUseCase,
		loginUseCase:  loginUseCase,
	}
}

func (h *AuthHandler) SignUp(c *gin.Context) {
	var req dto.SignUpRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := h.signUpUseCase.Execute(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, response)
}

func (h *AuthHandler) Login(c *gin.Context) {
	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := h.loginUseCase.Execute(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)
}
