package usecase

import (
	"github.com/yuhari7/article_service/internal/entity"
)

func (u *articleUsecase) GetAllArticles(limit, offset int, category, status, search string) ([]entity.Post, error) {
	return u.repo.FindAll(limit, offset, category, status, search)
}
