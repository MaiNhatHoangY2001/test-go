package handler

import (
	"test-go/internal/features/auth/dto"
	"test-go/internal/features/auth/usecase"
	"test-go/internal/shared/response"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type AuthHandler struct {
	signUpUseCase *usecase.SignUpUseCase
	loginUseCase  *usecase.LoginUseCase
	logger        *logrus.Logger
}

func NewAuthHandler(
	signUpUseCase *usecase.SignUpUseCase,
	loginUseCase *usecase.LoginUseCase,
	logger *logrus.Logger,
) *AuthHandler {
	return &AuthHandler{
		signUpUseCase: signUpUseCase,
		loginUseCase:  loginUseCase,
		logger:        logger,
	}
}

func (h *AuthHandler) SignUp(c *gin.Context) {
	requestID := c.GetString("X-Request-ID")

	var req dto.SignUpRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.WithFields(logrus.Fields{
			"request_id": requestID,
			"error":      err.Error(),
			"action":     "sign_up",
		}).Warn("invalid_request_body")
		response.BadRequest(c, "Invalid request body")
		return
	}

	h.logger.WithFields(logrus.Fields{
		"request_id": requestID,
		"email":      req.Email,
		"action":     "sign_up",
	}).Info("sign_up_started")

	result, err := h.signUpUseCase.Execute(c.Request.Context(), &req)
	if err != nil {
		h.logger.WithFields(logrus.Fields{
			"request_id": requestID,
			"email":      req.Email,
			"error":      err.Error(),
			"action":     "sign_up",
		}).Error("sign_up_failed")
		response.HandleError(c, h.logger, err)
		return
	}

	h.logger.WithFields(logrus.Fields{
		"request_id": requestID,
		"email":      req.Email,
		"action":     "sign_up",
	}).Info("sign_up_completed")

	response.Created(c, result)
}

func (h *AuthHandler) Login(c *gin.Context) {
	requestID := c.GetString("X-Request-ID")

	var req dto.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		h.logger.WithFields(logrus.Fields{
			"request_id": requestID,
			"error":      err.Error(),
			"action":     "login",
		}).Warn("invalid_request_body")
		response.BadRequest(c, "Invalid request body")
		return
	}

	h.logger.WithFields(logrus.Fields{
		"request_id": requestID,
		"email":      req.Email,
		"action":     "login",
	}).Info("login_started")

	result, err := h.loginUseCase.Execute(c.Request.Context(), &req)
	if err != nil {
		h.logger.WithFields(logrus.Fields{
			"request_id": requestID,
			"email":      req.Email,
			"error":      err.Error(),
			"action":     "login",
		}).Error("login_failed")
		response.HandleError(c, h.logger, err)
		return
	}

	h.logger.WithFields(logrus.Fields{
		"request_id": requestID,
		"email":      req.Email,
		"action":     "login",
	}).Info("login_completed")

	response.OK(c, result)
}
