package server

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/taufiksidiqmulyadinata15/go-backend-boilerplate/internal/config"
	"github.com/taufiksidiqmulyadinata15/go-backend-boilerplate/internal/handler"
	appMiddleware "github.com/taufiksidiqmulyadinata15/go-backend-boilerplate/internal/middleware"
	"github.com/taufiksidiqmulyadinata15/go-backend-boilerplate/internal/repository"
	"github.com/taufiksidiqmulyadinata15/go-backend-boilerplate/internal/service"
	"github.com/taufiksidiqmulyadinata15/go-backend-boilerplate/pkg/response"
	appValidator "github.com/taufiksidiqmulyadinata15/go-backend-boilerplate/pkg/validator"
	"gorm.io/gorm"
)

func NewServer(cfg *config.Config, db *gorm.DB) *echo.Echo {
	e := echo.New()
	e.Validator = appValidator.NewValidator()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	userRepo := repository.NewUserRepository(db)
	authService := service.NewAuthService(userRepo, cfg)
	authHandler := handler.NewAuthHandler(authService)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)
	e.GET("/", func(c echo.Context) error {
		return response.OK(c, "Go Backend Boilerplate is running", map[string]any{
			"service": cfg.AppName,
			"env":     cfg.AppEnv,
		})
	})

	e.GET("/health", func(c echo.Context) error {
		sqlDB, err := db.DB()
		if err != nil {
			return response.InternalServerError(c, "Database error", err.Error())
		}

		if err := sqlDB.Ping(); err != nil {
			return response.InternalServerError(c, "Database not connected", err.Error())
		}

		return response.Success(c, http.StatusOK, "Health check success", map[string]any{
			"status":   "ok",
			"database": "connected",
		})
	})

	api := e.Group("/api/v1")

	auth := api.Group("/auth")
	auth.POST("/register", authHandler.Register)
	auth.POST("/login", authHandler.Login)

	protected := api.Group("")
	protected.Use(appMiddleware.JWTMiddleware(cfg))

	protected.GET("/me", userHandler.Me)

	users := protected.Group("/users")
	users.Use(appMiddleware.AdminOnly)

	users.GET("", userHandler.GetUsers)
	users.GET("/:id", userHandler.GetUserByID)
	users.PUT("/:id", userHandler.UpdateUser)
	users.DELETE("/:id", userHandler.DeleteUser)

	return e
}
