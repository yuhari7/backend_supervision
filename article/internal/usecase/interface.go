package usecase

import (
	"github.com/yuhari7/article_service/internal/entity"
)

type ArticleUsecase interface {
	GetAllArticles(limit, offset int, category, status, search string) ([]entity.Post, error)
	CreateArticle(article entity.Post) (entity.Post, error)
	UpdateArticle(id uint, article entity.Post) (entity.Post, error)
	DeleteArticle(id uint) error
}
