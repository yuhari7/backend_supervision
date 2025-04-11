package user

import (
	"github.com/yuhari7/backend_supervision/internal/common/dto"
	"github.com/yuhari7/backend_supervision/internal/entity"
)

type UserUsecase interface {
	Register(input dto.CreateUserRequest) (*entity.User, error)
	Login(input dto.LoginRequest) (*entity.User, error)
	GetUserByID(id uint) (*entity.User, error)
	GetAllUsers(pagination dto.PaginationQuery) (*dto.PaginatedResponse[dto.UserResponse], error)
	UpdateUser(id uint, input dto.UpdateUserRequest) (*entity.User, error)
	ToggleUserActive(id uint, active bool) error
	DeleteUser(id uint) error
}
