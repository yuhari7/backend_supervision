package user

import "github.com/yuhari7/backend_supervision/internal/common/dto"

func (u *userUsecase) GetAllUsers(p dto.PaginationQuery) (*dto.PaginatedResponse[dto.UserResponse], error) {
	if p.Page <= 0 {
		p.Page = 1
	}
	if p.Limit <= 0 {
		p.Limit = 10
	}

	offset := (p.Page - 1) * p.Limit

	// optional search
	users, err := u.userRepo.FindWithPagination(p.Search, p.Limit, offset)
	if err != nil {
		return nil, err
	}

	total, err := u.userRepo.CountUsers(p.Search)
	if err != nil {
		return nil, err
	}

	// mapping to response
	var response []dto.UserResponse
	for _, u := range users {
		response = append(response, dto.UserResponse{
			ID:    u.ID,
			Name:  u.Name,
			Email: u.Email,
			Role:  u.RoleID,
		})
	}

	return &dto.PaginatedResponse[dto.UserResponse]{
		Data:       response,
		Page:       p.Page,
		Limit:      p.Limit,
		Total:      total,
		TotalPages: (total + p.Limit - 1) / p.Limit,
	}, nil
}
