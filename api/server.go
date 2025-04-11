package api

import (
	"github.com/labstack/echo/v4"

	controller "github.com/yuhari7/backend_supervision/api/controller/User"
	"github.com/yuhari7/backend_supervision/config"
	"github.com/yuhari7/backend_supervision/internal/repository"
	"github.com/yuhari7/backend_supervision/internal/usecase/user"
)

func NewServer() *echo.Echo {
	e := echo.New()

	// Dependency injection
	userRepo := repository.NewUserRepository(config.DB)
	userUsecase := user.NewUserUsecase(userRepo)

	// Register routes
	api := e.Group("/api")
	controller.RegisterUserRoutes(api, userUsecase)

	return e
}
