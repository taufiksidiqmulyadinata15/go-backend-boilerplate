package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/taufiksidiqmulyadinata15/go-backend-boilerplate/pkg/response"
)

func AdminOnly(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		role, ok := c.Get("role").(string)
		if !ok || role != "admin" {
			return response.Error(c, 403, "Forbidden", "admin access required")
		}

		return next(c)
	}
}
