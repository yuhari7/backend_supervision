package api

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/yuhari7/backend_supervision/article/api/controller"
	"github.com/yuhari7/backend_supervision/article/config"
	"github.com/yuhari7/backend_supervision/article/internal/repository"
	"github.com/yuhari7/backend_supervision/article/internal/usecase"
)

func NewServer() *echo.Echo {

	config.InitDB()

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
		AllowHeaders: []string{echo.HeaderContentType, echo.HeaderAuthorization},
	}))

	articleRepo := repository.NewArticleRepository()
	articleUsecase := usecase.NewArticleUsecase(articleRepo)
	articleController := controller.NewArticleController(articleUsecase)

	api := e.Group("/api")
	controller.RegisterArticleRoutes(api, articleController)

	log.Println("✅ Starting server on port 8001...")
	e.Logger.Fatal(e.Start(":8001"))

	return e
}
