package handler

import (
	"net/http"

	"github.com/taufiksidiqmulyadinata15/go-backend-boilerplate/internal/model"
	"github.com/taufiksidiqmulyadinata15/go-backend-boilerplate/internal/service"
	"github.com/taufiksidiqmulyadinata15/go-backend-boilerplate/pkg/response"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	UserService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{UserService: userService}
}

func (h *UserHandler) Me(c echo.Context) error {
	return response.OK(c, "Profile fetched successfully", map[string]any{
		"user_id": c.Get("user_id"),
		"email":   c.Get("email"),
		"role":    c.Get("role"),
	})
}

func (h *UserHandler) GetUsers(c echo.Context) error {
	users, err := h.UserService.GetUsers(
		c.QueryParam("page"),
		c.QueryParam("limit"),
		c.QueryParam("search"),
		c.QueryParam("role"),
	)
	if err != nil {
		return response.FromError(c, err)
	}

	return response.OK(c, "Users fetched successfully", users)
}
func (h *UserHandler) GetUserByID(c echo.Context) error {
	user, err := h.UserService.GetUserByID(c.Param("id"))
	if err != nil {
		return response.FromError(c, err)
	}

	return response.OK(c, "User fetched successfully", user)
}

func (h *UserHandler) UpdateUser(c echo.Context) error {
	var req model.UpdateUserRequest

	if err := c.Bind(&req); err != nil {
		return response.BadRequest(c, "Invalid request body", err.Error())
	}
	if err := c.Validate(&req); err != nil {
		return response.BadRequest(c, "Validation failed", err.Error())
	}

	user, err := h.UserService.UpdateUser(c.Param("id"), req.Name, req.Role)
	if err != nil {
		return response.FromError(c, err)
	}

	return response.OK(c, "User updated successfully", user)
}

func (h *UserHandler) DeleteUser(c echo.Context) error {
	currentUserID, ok := c.Get("user_id").(uint)
	if !ok {
		return response.Error(c, http.StatusUnauthorized, "Unauthorized", "invalid user token")
	}

	if err := h.UserService.DeleteUser(c.Param("id"), currentUserID); err != nil {
		return response.FromError(c, err)
	}

	return response.OK(c, "User deleted successfully", nil)
}
