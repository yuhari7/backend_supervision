package controller

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/yuhari7/backend_supervision/article/internal/common/dto"
	"github.com/yuhari7/backend_supervision/article/internal/usecase"
)

type ArticleController struct {
	ArticleUsecase usecase.ArticleUsecase
}

// NewArticleController creates a new instance of ArticleController
func NewArticleController(articleUsecase usecase.ArticleUsecase) *ArticleController {
	return &ArticleController{
		ArticleUsecase: articleUsecase,
	}
}

// Create handles the creation of a new article
func (c *ArticleController) Create(ctx echo.Context) error {
	var dto dto.CreateArticleRequest
	if err := ctx.Bind(&dto); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "invalid request"})
	}

	// Execute the create article usecase
	response, err := c.ArticleUsecase.CreateArticle(dto)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return ctx.JSON(http.StatusCreated, response)
}

// GetAll handles retrieving all articles
func (c *ArticleController) GetAll(ctx echo.Context) error {
	// Get pagination parameters from query
	limitStr := ctx.QueryParam("limit")
	offsetStr := ctx.QueryParam("offset")

	// Convert limit and offset to integers
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		limit = 10 // Set default limit to 10 if no value is provided or invalid
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		offset = 0 // Set default offset to 0 if no value is provided or invalid
	}

	// Execute the find all articles usecase with pagination
	articles, err := c.ArticleUsecase.FindAllArticles(limit, offset)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	// Return the paginated list of articles
	return ctx.JSON(http.StatusOK, articles)
}

func (c *ArticleController) FindByID(ctx echo.Context) error {
	// Get the article ID from the URL parameter
	idStr := ctx.Param("id")

	// Convert the string ID to uint
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "invalid article ID"})
	}

	// Execute the find article usecase
	article, err := c.ArticleUsecase.FindByID(uint(id))
	if err != nil {
		return ctx.JSON(http.StatusNotFound, echo.Map{"error": err.Error()})
	}

	// Return the article details
	return ctx.JSON(http.StatusOK, article)
}

func (c *ArticleController) Update(ctx echo.Context) error {
	var dto dto.UpdateArticleRequest
	if err := ctx.Bind(&dto); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "invalid request"})
	}

	// Get the article ID from the URL parameter (as a string)
	idStr := ctx.Param("id")

	// Convert the string ID to uint
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "invalid article ID"})
	}

	// Set the ID for the update request
	dto.ID = uint(id)

	// Execute the update article usecase
	article, err := c.ArticleUsecase.UpdateArticle(dto)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, echo.Map{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, article) // Return the updated article
}

func (c *ArticleController) SoftDelete(ctx echo.Context) error {
	// Get the article ID from the URL parameter
	idStr := ctx.Param("id")

	// Convert the string ID to uint
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "invalid article ID"})
	}

	// Set the ID for the soft delete request
	dto := dto.SoftDeleteArticleDTO{
		ID: uint(id),
	}

	// Execute the soft delete article usecase
	article, err := c.ArticleUsecase.SoftDeleteArticle(dto)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, article) // Return the updated article
}

// Delete handles the permanent deletion of an article
func (c *ArticleController) Delete(ctx echo.Context) error {
	// Get the article ID from the URL parameter
	idStr := ctx.Param("id")

	// Convert the string ID to uint
	id, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "invalid article ID"})
	}

	// Set the ID for the delete request
	dto := dto.SoftDeleteArticleDTO{
		ID: uint(id),
	}

	// Execute the permanent delete article usecase
	err = c.ArticleUsecase.DeleteArticle(dto)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, echo.Map{"message": "Article deleted permanently"})
}
