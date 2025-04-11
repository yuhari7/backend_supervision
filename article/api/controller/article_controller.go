package controller

import (
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/yuhari7/article_service/internal/entity"
	"github.com/yuhari7/article_service/internal/usecase"
)

type ArticleController struct {
	Usecase   usecase.ArticleUsecase
	Validator *validator.Validate
}

func NewArticleController(u usecase.ArticleUsecase) *ArticleController {
	return &ArticleController{Usecase: u, Validator: validator.New()}
}

func (h *ArticleController) validatePost(article *entity.Post) error {
	return h.Validator.Struct(article)
}

func (h *ArticleController) GetAll(c echo.Context) error {
	// Get pagination parameters from query params
	page, _ := strconv.Atoi(c.QueryParam("page"))
	limit, _ := strconv.Atoi(c.QueryParam("limit"))
	category := c.QueryParam("category")
	status := c.QueryParam("status")
	search := c.QueryParam("search")

	// Default to page 1 and limit 10 if not provided
	if page == 0 {
		page = 1
	}
	if limit == 0 {
		limit = 10
	}

	offset := (page - 1) * limit

	// Fetch posts with pagination, filtering, and search
	posts, err := h.Usecase.GetAllArticles(limit, offset, category, status, search)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	// Return the posts with pagination info
	return c.JSON(http.StatusOK, echo.Map{
		"posts":    posts,
		"page":     page,
		"limit":    limit,
		"category": category,
		"status":   status,
		"search":   search,
	})
}

func (h *ArticleController) Create(c echo.Context) error {
	var article entity.Post
	if err := c.Bind(&article); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid request"})
	}

	//validasi
	if err := h.validatePost(&article); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": formatValidationErrors(err)})
	}

	createdArticle, err := h.Usecase.CreateArticle(article)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, echo.Map{"article": createdArticle})
}

func (h *ArticleController) Update(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	var article entity.Post
	if err := c.Bind(&article); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "invalid request"})
	}

	// validasi
	if err := h.validatePost(&article); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": formatValidationErrors(err)})
	}

	updatedArticle, err := h.Usecase.UpdateArticle(uint(id), article)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, updatedArticle)
}

func (h *ArticleController) Delete(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	err := h.Usecase.DeleteArticle(uint(id))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": err.Error()})
	}
	return c.JSON(http.StatusOK, echo.Map{"message": "Article deleted successfully"})
}

func formatValidationErrors(err error) string {
	var errorMessage string
	for _, e := range err.(validator.ValidationErrors) {
		switch e.Tag() {
		case "required":
			errorMessage += e.Field() + " is required. "
		case "min":
			// Custom error messages for min length validation
			if e.Field() == "Title" {
				errorMessage += "Title must be at least 20 characters long. "
			} else if e.Field() == "Content" {
				errorMessage += "Content must be at least 200 characters long. "
			} else if e.Field() == "Category" {
				errorMessage += "Category must be at least 3 characters long. "
			}
		}
	}
	return errorMessage
}
