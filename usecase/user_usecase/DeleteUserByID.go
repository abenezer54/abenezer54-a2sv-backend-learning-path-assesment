package user_usecase

import (
	"context"
	"errors"
)

func (u *userUsecase) DeleteUserByID(ctx context.Context, id string) error {
	user, err := u.userRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	if user == nil {
		return errors.New("user not found")
	}

	err = u.userRepo.Delete(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
