package user_usecase

import (
	"context"
	"errors"
)

func (u *userUsecase) VerifyEmail(ctx context.Context, token string) error {
	// Validate the token
	email, err := u.VerifyTokenRepo.ValidateVerifyToken(ctx, token)
	if err != nil {
		return errors.New("invalid or expired token")
	}

	// Fetch the user by email
	user, err := u.userRepo.GetByEmail(ctx, email)
	if err != nil {
		return err
	}

	if user == nil {
		return errors.New("user not found")
	}

	// Invalidate the verification token
	err = u.VerifyTokenRepo.InvalidateVerifyToken(ctx, token)
	if err != nil {
		return err
	}

	return nil
}
