package dto

// CreateArticleRequest represents the data required to create an article
type CreateArticleRequest struct {
	Title    string `json:"title" validate:"required,min=20"`
	Content  string `json:"content" validate:"required,min=200"`
	Category string `json:"category" validate:"required,min=3"`
	Status   string `json:"status"`
}

// CreateArticleResponse represents the response data after creating an article
type CreateArticleResponse struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Category  string `json:"category"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_date"`
	UpdatedAt string `json:"updated_date"`
}

// UpdateArticleRequest represents the data required to update an article
type UpdateArticleRequest struct {
	ID       uint   `json:"id"`
	Title    string `json:"title" validate:"required,min=20"`
	Content  string `json:"content" validate:"required,min=200"`
	Category string `json:"category" validate:"required,min=3"`
	Status   string `json:"status"`
}

// UpdateArticleResponse represents the response data after updating an article
type UpdateArticleResponse struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Category  string `json:"category"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_date"`
	UpdatedAt string `json:"updated_date"`
}

// SoftDeleteArticleDTO represents the data required to soft delete an article
type SoftDeleteArticleDTO struct {
	ID uint `json:"id"`
}

type ArticleResponse struct {
	ID        uint   `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Category  string `json:"category"`
	Status    string `json:"status"`
	CreatedAt string `json:"created_date"`
	UpdatedAt string `json:"updated_date"`
}
