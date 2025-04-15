package repository

import (
	"time"

	"github.com/yuhari7/backend_supervision/article/config"
	"github.com/yuhari7/backend_supervision/article/internal/entity"
	"gorm.io/gorm"
)

// ArticleRepository defines the methods for interacting with the database
type ArticleRepository interface {
	Create(article *entity.Article) error
	FindAll() ([]entity.Article, error)
	FindByID(id uint) (*entity.Article, error)
	Update(article *entity.Article) error
	Delete(id uint) error
	SoftDelete(id uint) error
	FindWithPagination(limit, offset int, articles *[]entity.Article) error // Add pagination method
}

type articleRepository struct{}

// NewArticleRepository creates a new instance of ArticleRepository
func NewArticleRepository() ArticleRepository {
	return &articleRepository{}
}

// Create inserts a new article into the database
func (r *articleRepository) Create(article *entity.Article) error {
	return config.DB.Create(article).Error
}

// FindAll returns all articles from the database
func (r *articleRepository) FindAll() ([]entity.Article, error) {
	var articles []entity.Article
	err := config.DB.Find(&articles).Error
	return articles, err
}

// FindByID finds an article by its ID
func (r *articleRepository) FindByID(id uint) (*entity.Article, error) {
	var article entity.Article
	err := config.DB.First(&article, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil // If the article is not found, return nil
		}
		return nil, err // Return other errors
	}
	return &article, nil
}

// Update updates an article in the database
func (r *articleRepository) Update(article *entity.Article) error {
	return config.DB.Save(article).Error
}

// Delete deletes an article by its ID
func (r *articleRepository) Delete(id uint) error {
	return config.DB.Delete(&entity.Article{}, id).Error
}

// SoftDelete sets the deleted_at timestamp to implement soft delete
func (r *articleRepository) SoftDelete(id uint) error {
	return config.DB.Model(&entity.Article{}).Where("id = ?", id).Update("deleted_at", time.Now()).Error
}

func (r *articleRepository) FindWithPagination(limit, offset int, articles *[]entity.Article) error {
	return config.DB.Limit(limit).Offset(offset).Find(articles).Error
}
