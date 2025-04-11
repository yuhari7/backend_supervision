package repository

import (
	"errors"

	"github.com/yuhari7/article_service/config"
	"github.com/yuhari7/article_service/internal/entity"
)

type ArticleRepository interface {
	FindAll(limit, offset int, category, status, search string) ([]entity.Post, error)
	Create(article *entity.Post) error
	Delete(id uint) error
	FindByID(id uint, post *entity.Post) error
	Update(post *entity.Post) error
}

type articleRepository struct{}

func NewArticleRepository() ArticleRepository {
	return &articleRepository{}
}

func (r *articleRepository) FindAll(limit, offset int, category, status, search string) ([]entity.Post, error) {
	var articles []entity.Post
	query := config.DB.Limit(limit).Offset(offset)

	if category != "" {
		query = query.Where("category = ?", category)
	}
	if status != "" {
		query = query.Where("status = ?", status)
	}

	if search != "" {
		query = query.Where("title LIKE ? OR content LIKE ?", "%"+search+"%", "%"+search+"%")
	}

	err := query.Find(&articles).Error
	return articles, err
}

// FindByID finds a post by ID
func (r *articleRepository) FindByID(id uint, post *entity.Post) error {
	if err := config.DB.First(post, id).Error; err != nil {
		return errors.New("post not found")
	}
	return nil
}

func (r *articleRepository) Update(post *entity.Post) error {
	if err := config.DB.Save(post).Error; err != nil {
		return err
	}
	return nil
}

func (r *articleRepository) Create(article *entity.Post) error {
	return config.DB.Create(article).Error
}

func (r *articleRepository) Delete(id uint) error {
	if err := config.DB.Delete(&entity.Post{}, id).Error; err != nil {
		return err
	}
	return nil
}
