package controller

import (
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/yuhari7/backend_supervision/article/internal/common/dto"
	"github.com/yuhari7/backend_supervision/article/internal/usecase"
)

type ArticleController struct {
	ArticleUsecase usecase.ArticleUsecase
	Validator      *validator.Validate
}

// NewArticleController creates a new instance of ArticleController
func NewArticleController(articleUsecase usecase.ArticleUsecase) *ArticleController {
	return &ArticleController{
		ArticleUsecase: articleUsecase,
		Validator:      validator.New(),
	}
}

// Create handles the creation of a new article
func (c *ArticleController) Create(ctx echo.Context) error {
	var request dto.CreateArticleRequest

	// Bind the incoming request body to the DTO
	if err := ctx.Bind(&request); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "invalid request"})
	}

	// Validate the request data
	if err := c.Validator.Struct(&request); err != nil {
		errorMessages := make(map[string]string)
		for _, err := range err.(validator.ValidationErrors) {
			switch err.Field() {
			case "Title":
				errorMessages["title"] = "Title Kurang dari 20 Character"
			case "Content":
				errorMessages["content"] = "Description Minimal 200 Character"
			case "Category":
				errorMessages["category"] = "Category minimal 3 Character"
			default:
				errorMessages[err.Field()] = err.Error()
			}
		}
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": errorMessages})
	}

	// Call the usecase to create the article
	article, err := c.ArticleUsecase.CreateArticle(request)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return ctx.JSON(http.StatusCreated, article)
}

// GetAll handles retrieving all articles
func (c *ArticleController) GetAll(ctx echo.Context) error {
	// Get pagination parameters from query
	limitStr := ctx.QueryParam("limit")
	offsetStr := ctx.QueryParam("offset")

	// Convert limit and offset to integers
	limit, err := strconv.Atoi(limitStr)
	if err != nil {
		limit = 10
	}

	offset, err := strconv.Atoi(offsetStr)
	if err != nil {
		offset = 0
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
	var request dto.UpdateArticleRequest

	if err := ctx.Bind(&request); err != nil {
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": "invalid request"})
	}

	// Validate the request data
	if err := c.Validator.Struct(&request); err != nil {
		errorMessages := make(map[string]string)

		for _, err := range err.(validator.ValidationErrors) {
			switch err.Field() {
			case "Title":
				errorMessages["title"] = "Title Kurang dari 20 Character"
			case "Content":
				errorMessages["content"] = "Description Minimal 200 Character"
			case "Category":
				errorMessages["category"] = "Category minimal 3 Character"
			default:
				errorMessages[err.Field()] = err.Error()
			}
		}

		// Return validation errors
		return ctx.JSON(http.StatusBadRequest, echo.Map{"error": errorMessages})
	}

	// Execute the update article usecase
	article, err := c.ArticleUsecase.UpdateArticle(request)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
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

func (c *ArticleController) Search(ctx echo.Context) error {
	query := ctx.QueryParam("q")
	limitStr := ctx.QueryParam("limit")
	offsetStr := ctx.QueryParam("offset")

	limit := 10
	offset := 0

	if limitStr != "" {
		if parsedLimit, err := strconv.Atoi(limitStr); err == nil {
			limit = parsedLimit
		}
	}
	if offsetStr != "" {
		if parsedOffset, err := strconv.Atoi(offsetStr); err == nil {
			offset = parsedOffset
		}
	}

	if query == "" {
		return ctx.JSON(http.StatusBadRequest, map[string]string{"error": "Query parameter 'q' is required"})
	}

	articles, err := c.ArticleUsecase.SearchArticles(query, limit, offset)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return ctx.JSON(http.StatusOK, articles)
}
