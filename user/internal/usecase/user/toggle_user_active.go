package user

import "errors"

func (u *userUsecase) ToggleUserActive(id uint, active bool) error {
	user, err := u.userRepo.FindByID(id)
	if err != nil {
		return err
	}
	if user == nil {
		return errors.New("user not found")
	}

	user.IsActive = active
	return u.userRepo.Update(user)
}
