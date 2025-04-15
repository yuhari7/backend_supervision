package api

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	controller "github.com/yuhari7/backend_supervision/api/controller/User"
	"github.com/yuhari7/backend_supervision/config"
	"github.com/yuhari7/backend_supervision/internal/repository"
	"github.com/yuhari7/backend_supervision/internal/usecase/user"
)

func NewServer() *echo.Echo {
	e := echo.New()

	// Set up CORS middleware
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	// Dependency injection
	userRepo := repository.NewUserRepository(config.DB)
	userUsecase := user.NewUserUsecase(userRepo)

	// Register routes
	api := e.Group("/api")
	controller.RegisterUserRoutes(api, userUsecase)

	return e
}
