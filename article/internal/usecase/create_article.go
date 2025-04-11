package usecase

import (
	"github.com/yuhari7/article_service/internal/entity"
	"github.com/yuhari7/article_service/internal/repository"
)

type articleUsecase struct {
	repo repository.ArticleRepository
}

// Constructor for creating a new ArticleUsecase
func NewArticleUsecase(r repository.ArticleRepository) ArticleUsecase {
	return &articleUsecase{repo: r}
}

func (u *articleUsecase) CreateArticle(article entity.Post) (entity.Post, error) {
	err := u.repo.Create(&article)
	return article, err
}
