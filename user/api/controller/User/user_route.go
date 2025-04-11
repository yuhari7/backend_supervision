package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/yuhari7/backend_supervision/api/middleware"
	"github.com/yuhari7/backend_supervision/internal/usecase/user"
)

func RegisterUserRoutes(e *echo.Group, usecase user.UserUsecase) {
	handler := NewUserController(usecase)

	e.POST("/register", handler.Register)
	e.POST("/login", handler.Login)
	e.POST("/refresh", handler.RefreshToken)

	protected := e.Group("/users", middleware.AuthMiddleware, middleware.AdminOnlyMiddleware)
	protected.GET("", handler.GetAllUsers)
	protected.GET("/:id", handler.GetUserByID)
	protected.POST("", handler.CreateUser)
	protected.PUT("/:id", handler.UpdateUser)
	protected.DELETE("/:id", handler.DeleteUser)
	protected.PUT("/:id/deactivate", handler.DeactivateUser)
	protected.PUT("/:id/activate", handler.ActivateUser)

}
