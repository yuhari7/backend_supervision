package user

import "github.com/yuhari7/backend_supervision/internal/entity"

func (u *userUsecase) GetUserByID(id uint) (*entity.User, error) {
	return u.userRepo.FindByID(id)
}
