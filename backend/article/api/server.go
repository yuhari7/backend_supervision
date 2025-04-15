package api

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/yuhari7/backend_supervision/article/api/controller"
	"github.com/yuhari7/backend_supervision/article/config"
	"github.com/yuhari7/backend_supervision/article/internal/repository"
	"github.com/yuhari7/backend_supervision/article/internal/usecase"
)

func NewServer() *echo.Echo {

	config.InitDB()

	e := echo.New()

	articleRepo := repository.NewArticleRepository()
	articleUsecase := usecase.NewArticleUsecase(articleRepo)
	articleController := controller.NewArticleController(articleUsecase)

	api := e.Group("/api")
	controller.RegisterArticleRoutes(api, articleController)

	log.Println("✅ Starting server on port 8001...")
	e.Logger.Fatal(e.Start(":8001"))

	return e
}
