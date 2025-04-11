package api

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/yuhari7/article_service/api/controller"
)

// NewServer initializes Echo instance and registers routes
func NewServer() *echo.Echo {
	e := echo.New()

	// Set up CORS middleware
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:3000"},                                                       // Allow from frontend
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},                                     // Add OPTIONS for preflight requests
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, "Authorization"}, // Add 'Authorization' here
		AllowCredentials: true,                                                                                    // Allow credentials (cookies, etc.)
	}))

	// Register routes for articles
	controller.RegisterArticleRoutes(e)

	return e
}
