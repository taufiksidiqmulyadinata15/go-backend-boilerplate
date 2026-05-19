package middleware

import (
	"strings"

	"github.com/taufiksidiqmulyadinata15/go-backend-boilerplate/internal/config"
	jwthelper "github.com/taufiksidiqmulyadinata15/go-backend-boilerplate/pkg/jwt"
	"github.com/taufiksidiqmulyadinata15/go-backend-boilerplate/pkg/response"

	"github.com/labstack/echo/v4"
)

func JWTMiddleware(cfg *config.Config) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			authHeader := c.Request().Header.Get("Authorization")

			if authHeader == "" {
				return response.Error(c, 401, "Unauthorized", "missing authorization header")
			}

			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
				return response.Error(c, 401, "Unauthorized", "invalid authorization header")
			}

			claims, err := jwthelper.ValidateToken(parts[1], cfg.JWTSecret)
			if err != nil {
				return response.Error(c, 401, "Unauthorized", "invalid or expired token")
			}

			c.Set("user_id", claims.UserID)
			c.Set("email", claims.Email)
			c.Set("role", claims.Role)

			return next(c)
		}
	}
}
