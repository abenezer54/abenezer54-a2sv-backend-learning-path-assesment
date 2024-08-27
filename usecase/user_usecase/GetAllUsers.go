package user_usecase

import (
	"context"
	"loan-api/domain"
)

func (u *userUsecase) GetAllUsers(ctx context.Context) ([]domain.User, error) {
	users, err := u.userRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return users, nil
}
