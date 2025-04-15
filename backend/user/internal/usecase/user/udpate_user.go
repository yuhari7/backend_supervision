package user

import (
	"errors"
	"strings"

	"github.com/yuhari7/backend_supervision/internal/common/dto"
	"github.com/yuhari7/backend_supervision/internal/entity"
	"golang.org/x/crypto/bcrypt"
)

type UpdateInput struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password,omitempty"` // optional
	RoleID   uint   `json:"role_id"`
}

func (u *userUsecase) UpdateUser(id uint, input dto.UpdateUserRequest) (*entity.User, error) {
	user, err := u.userRepo.FindByID(id)
	if err != nil {
		return nil, err
	}

	user.Name = input.Name
	user.Email = input.Email
	user.RoleID = input.RoleID

	if strings.TrimSpace(input.Password) != "" {
		hashed, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
		if err != nil {
			return nil, errors.New("failed to hash password")
		}
		user.Password = string(hashed)
	}

	err = u.userRepo.Update(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}
