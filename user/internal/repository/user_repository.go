package repository

import (
	"errors"

	"github.com/yuhari7/backend_supervision/internal/entity"
	"gorm.io/gorm"
)

// Interface
type UserRepository interface {
	Create(user *entity.User) error
	FindByEmail(email string) (*entity.User, error)
	FindByID(id uint) (*entity.User, error)
	FindAll() ([]entity.User, error)
	Delete(id uint) error
	Update(user *entity.User) error
	FindWithPagination(search string, limit, offset int) ([]entity.User, error)
	CountUsers(search string) (int, error)
}

// Implementation
type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user *entity.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) FindByEmail(email string) (*entity.User, error) {
	var user entity.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &user, err
}

func (r *userRepository) FindByID(id uint) (*entity.User, error) {
	var user entity.User
	err := r.db.First(&user, id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return &user, err
}

func (r *userRepository) FindAll() ([]entity.User, error) {
	var users []entity.User
	err := r.db.Find(&users).Error
	return users, err
}

// func (r *userRepository) Delete(id uint) error {
// 	return r.db.Delete(&entity.User{}, id).Error
// }

func (r *userRepository) Delete(id uint) error {
	return r.db.Delete(&entity.User{}, id).Error
}

func (r *userRepository) Update(user *entity.User) error {
	return r.db.Save(user).Error
}

func (r *userRepository) FindWithPagination(search string, limit, offset int) ([]entity.User, error) {
	var users []entity.User
	query := r.db.Model(&entity.User{})

	if search != "" {
		search = "%" + search + "%"
		query = query.Where("name ILIKE ? OR email ILIKE ?", search, search)
	}

	err := query.Limit(limit).Offset(offset).Find(&users).Error
	return users, err
}

func (r *userRepository) CountUsers(search string) (int, error) {
	var count int64
	query := r.db.Model(&entity.User{})

	if search != "" {
		search = "%" + search + "%"
		query = query.Where("name ILIKE ? OR email ILIKE ?", search, search)
	}

	err := query.Count(&count).Error
	return int(count), err
}
