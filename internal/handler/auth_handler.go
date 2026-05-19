package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/taufiksidiqmulyadinata15/go-backend-boilerplate/internal/model"
	"github.com/taufiksidiqmulyadinata15/go-backend-boilerplate/internal/service"
	"github.com/taufiksidiqmulyadinata15/go-backend-boilerplate/pkg/response"
)

type AuthHandler struct {
	AuthService *service.AuthService
}

func NewAuthHandler(authService *service.AuthService) *AuthHandler {
	return &AuthHandler{AuthService: authService}
}

func (h *AuthHandler) Register(c echo.Context) error {
	var req model.RegisterRequest

	if err := c.Bind(&req); err != nil {
		return response.BadRequest(c, "Invalid request body", err.Error())
	}
	if err := c.Validate(&req); err != nil {
		return response.BadRequest(c, "Validation failed", err.Error())
	}
	user, err := h.AuthService.Register(req)
	if err != nil {
		return response.BadRequest(c, err.Error(), nil)
	}

	return response.Created(c, "User registered successfully", user)
}
func (h *AuthHandler) Login(c echo.Context) error {
	var req model.LoginRequest

	if err := c.Bind(&req); err != nil {
		return response.BadRequest(c, "Invalid request body", err.Error())
	}
	if err := c.Validate(&req); err != nil {
		return response.BadRequest(c, "Validation failed", err.Error())
	}
	result, err := h.AuthService.Login(req)
	if err != nil {
		return response.BadRequest(c, err.Error(), nil)
	}

	return response.OK(c, "Login successfully", result)
}
