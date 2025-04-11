package user

func (u *userUsecase) DeleteUser(id uint) error {
	return u.userRepo.Delete(id)
}
