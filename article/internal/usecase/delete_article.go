package usecase

func (u *articleUsecase) DeleteArticle(id uint) error {
	return u.repo.Delete(id)
}
