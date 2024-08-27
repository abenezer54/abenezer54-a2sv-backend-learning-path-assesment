package user_usecase

import (
	"context"
	"loan-api/domain"
)

func (u *userUsecase) GetUserByID(ctx context.Context, id string) (*domain.User, error) {
	return u.userRepo.GetByID(ctx, id)
}
