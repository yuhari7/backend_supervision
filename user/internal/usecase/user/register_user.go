package user

import (
	"errors"

	"github.com/yuhari7/backend_supervision/internal/common/dto"
	"github.com/yuhari7/backend_supervision/internal/entity"
	"github.com/yuhari7/backend_supervision/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type RegisterInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	RoleID   uint   `json:"role_id"`
}

type userUsecase struct {
	userRepo repository.UserRepository
}

func NewUserUsecase(repo repository.UserRepository) UserUsecase {
	return &userUsecase{userRepo: repo}
}

func (u *userUsecase) Register(input dto.CreateUserRequest) (*entity.User, error) {
	// Cek email
	existing, _ := u.userRepo.FindByEmail(input.Email)
	if existing != nil {
		return nil, errors.New("email already registered")
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, errors.New("failed to hash password")
	}

	user := &entity.User{
		Name:     input.Name,
		Email:    input.Email,
		Password: string(hashed),
		RoleID:   input.RoleID,
	}

	err = u.userRepo.Create(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

// func (u *userUsecase) Register(input RegisterInput) (*entity.User, error) {
// 	// Validasi awal
// 	if strings.TrimSpace(input.Name) == "" || strings.TrimSpace(input.Email) == "" || strings.TrimSpace(input.Password) == "" {
// 		return nil, errors.New("name, email, and password are required")
// 	}

// 	// Cek email sudah digunakan?
// 	existingUser, _ := u.userRepo.FindByEmail(input.Email)
// 	if existingUser != nil {
// 		return nil, errors.New("email already registered")
// 	}

// 	// Hash password
// 	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
// 	if err != nil {
// 		return nil, errors.New("failed to hash password")
// 	}

// 	// Buat user
// 	user := &entity.User{
// 		Name:     input.Name,
// 		Email:    input.Email,
// 		Password: string(hashedPassword),
// 		RoleID:   input.RoleID,
// 	}

// 	err = u.userRepo.Create(user)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return user, nil
// }
