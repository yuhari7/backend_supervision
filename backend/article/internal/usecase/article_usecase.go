package usecase

import (
	"errors"

	"github.com/yuhari7/backend_supervision/article/internal/common/dto"
	"github.com/yuhari7/backend_supervision/article/internal/entity"
	"github.com/yuhari7/backend_supervision/article/internal/repository"
)

// ArticleUsecase defines the methods for interacting with articles
type ArticleUsecase interface {
	CreateArticle(dto dto.CreateArticleRequest) (entity.Article, error)
	UpdateArticle(dto dto.UpdateArticleRequest) (entity.Article, error)
	SoftDeleteArticle(dto dto.SoftDeleteArticleDTO) (entity.Article, error)
	DeleteArticle(dto dto.SoftDeleteArticleDTO) error
	FindByID(id uint) (dto.ArticleResponse, error)
	FindAllArticles(limit, offset int) ([]dto.ArticleResponse, error)
}

type articleUsecase struct {
	repo repository.ArticleRepository
}

// NewArticleUsecase creates a new instance of ArticleUsecase
func NewArticleUsecase(r repository.ArticleRepository) ArticleUsecase {
	return &articleUsecase{repo: r}
}

// FindAllArticles retrieves all articles from the repository
func (u *articleUsecase) FindAllArticles(limit, offset int) ([]dto.ArticleResponse, error) {
	var articles []entity.Article
	// Fetch articles with limit and offset
	err := u.repo.FindWithPagination(limit, offset, &articles)
	if err != nil {
		return nil, err
	}

	// Prepare response DTOs
	var responses []dto.ArticleResponse
	for _, article := range articles {
		response := dto.ArticleResponse{
			ID:        article.ID,
			Title:     article.Title,
			Content:   article.Content,
			Category:  article.Category,
			Status:    article.Status,
			CreatedAt: article.CreatedDate.Format("2006-01-02 15:04:05"),
			UpdatedAt: article.UpdatedDate.Format("2006-01-02 15:04:05"),
		}
		responses = append(responses, response)
	}

	return responses, nil
}

func (u *articleUsecase) FindByID(id uint) (dto.ArticleResponse, error) {
	article, err := u.repo.FindByID(id)
	if err != nil {
		return dto.ArticleResponse{}, errors.New("article not found")
	}

	// Convert the article entity to the response DTO
	response := dto.ArticleResponse{
		ID:        article.ID,
		Title:     article.Title,
		Content:   article.Content,
		Category:  article.Category,
		Status:    article.Status,
		CreatedAt: article.CreatedDate.Format("2006-01-02 15:04:05"),
		UpdatedAt: article.UpdatedDate.Format("2006-01-02 15:04:05"),
	}

	return response, nil
}

func (u *articleUsecase) CreateArticle(dto dto.CreateArticleRequest) (entity.Article, error) {
	article := entity.Article{
		Title:    dto.Title,
		Content:  dto.Content,
		Category: dto.Category,
		Status:   dto.Status,
	}

	// Save article to the repository (database)
	err := u.repo.Create(&article)
	if err != nil {
		return entity.Article{}, err
	}

	// Return the article entity directly
	return article, nil
}

func (u *articleUsecase) UpdateArticle(dto dto.UpdateArticleRequest) (entity.Article, error) {
	// Find the existing article by ID
	article, err := u.repo.FindByID(dto.ID)
	if err != nil {
		return entity.Article{}, errors.New("article not found")
	}

	// Update the article fields with the new data
	article.Title = dto.Title
	article.Content = dto.Content
	article.Category = dto.Category
	article.Status = dto.Status

	// Save the updated article to the repository (database)
	err = u.repo.Update(article) // Pass the pointer directly here
	if err != nil {
		return entity.Article{}, err
	}

	// Return the updated article entity
	return *article, nil // Return the dereferenced article (value)
}

func (u *articleUsecase) SoftDeleteArticle(dto dto.SoftDeleteArticleDTO) (entity.Article, error) {
	article, err := u.repo.FindByID(dto.ID)
	if err != nil {
		return entity.Article{}, errors.New("article not found")
	}

	// Set the status to "Trash"
	article.Status = "Trash"

	// Save the updated article
	err = u.repo.Update(article)
	if err != nil {
		return entity.Article{}, err
	}

	// Return the updated article
	return *article, nil
}

// DeleteArticle permanently removes an article if its status is Trash
func (u *articleUsecase) DeleteArticle(dto dto.SoftDeleteArticleDTO) error {
	article, err := u.repo.FindByID(dto.ID)
	if err != nil {
		return errors.New("article not found")
	}

	// Check if the article status is Trash (can be permanently deleted)
	if article.Status != "Trash" {
		return errors.New("only articles in trash can be permanently deleted")
	}

	// Permanently delete the article from the repository
	err = u.repo.Delete(dto.ID)
	if err != nil {
		return err
	}

	return nil
}
