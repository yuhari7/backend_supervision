package controller

import (
	"github.com/labstack/echo/v4"
)

// RegisterArticleRoutes sets up the routes for article-related endpoints
func RegisterArticleRoutes(e *echo.Group, controller *ArticleController) {
	e.POST("/articles", controller.Create)
	e.PUT("/articles/:id", controller.Update)
	e.GET("/articles", controller.GetAll)
	e.GET("/articles/:id", controller.FindByID)
	e.PUT("/articles/:id/trash", controller.SoftDelete)
	e.DELETE("/articles/:id", controller.Delete)
}
