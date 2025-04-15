package user

import (
	"errors"

	"github.com/yuhari7/backend_supervision/internal/common/dto"
	"github.com/yuhari7/backend_supervision/internal/entity"
	"golang.org/x/crypto/bcrypt"
)

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *userUsecase) Login(input dto.LoginRequest) (*entity.User, error) {
	user, err := u.userRepo.FindByEmail(input.Email)
	if err != nil {
		return nil, errors.New("failed to find user")
	}
	if user == nil {
		return nil, errors.New("user not found")
	}

	if !user.IsActive {
		return nil, errors.New("user is inactive")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	return user, nil
}
