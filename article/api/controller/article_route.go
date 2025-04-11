package controller

import (
	"github.com/labstack/echo/v4"
	"github.com/yuhari7/article_service/internal/repository"
	"github.com/yuhari7/article_service/internal/usecase"
)

// RegisterArticleRoutes function will register all the routes for articles
func RegisterArticleRoutes(e *echo.Echo) {
	// Initialize repositories
	articleRepo := repository.NewArticleRepository()

	// Initialize usecases
	articleUsecase := usecase.NewArticleUsecase(articleRepo)

	// Initialize controller
	articleController := NewArticleController(articleUsecase)

	// Register routes under the /api group
	api := e.Group("/api")
	api.GET("/articles", articleController.GetAll)
	api.POST("/articles", articleController.Create)
	api.PUT("/articles/:id", articleController.Update)
	api.DELETE("/articles/:id", articleController.Delete)
}
