package usecase

import (
	"errors"

	"github.com/yuhari7/article_service/internal/entity"
)

// Update the article
func (u *articleUsecase) UpdateArticle(id uint, article entity.Post) (entity.Post, error) {
	// Check if article exists
	var existingArticle entity.Post
	if err := u.repo.FindByID(id, &existingArticle); err != nil {
		return entity.Post{}, errors.New("Article not found")
	}

	// Update the article fields
	existingArticle.Title = article.Title
	existingArticle.Content = article.Content
	existingArticle.Category = article.Category
	existingArticle.Status = article.Status
	existingArticle.UpdatedDate = article.UpdatedDate

	// Save the updated article
	if err := u.repo.Update(&existingArticle); err != nil {
		return entity.Post{}, err
	}

	return existingArticle, nil
}
