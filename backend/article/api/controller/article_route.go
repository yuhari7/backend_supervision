package controller

import (
	"github.com/labstack/echo/v4"
)

// RegisterArticleRoutes sets up the routes for article-related endpoints
func RegisterArticleRoutes(e *echo.Group, controller *ArticleController) {
	// e.GET("/articles", controller.GetAll)

	articleGroup := e.Group("/articles")

	articleGroup.GET("", controller.GetAll)

	articleGroup.GET("/search", controller.Search)
	articleGroup.GET("/:id", controller.FindByID)

	articleGroup.POST("", controller.Create)

	articleGroup.PUT("/:id", controller.Update)
	articleGroup.PUT("/:id/trash", controller.SoftDelete)

	articleGroup.DELETE("/:id", controller.Delete)
}
