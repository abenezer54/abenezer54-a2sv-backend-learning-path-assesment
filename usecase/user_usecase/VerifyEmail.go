package user_usecase

import (
	"context"
	"errors"
	"loan-api/domain"
	"log"
)

func (u *userUsecase) VerifyEmail(ctx context.Context, token string) error {
	// Validate the token
	_, err := u.VerifyTokenRepo.ValidateVerifyToken(ctx, token)
	if err != nil {
		return errors.New("invalid or expired token")
	}

	// Check if user already exists in the database
	// user, err := u.userRepo.GetByEmail(ctx, email)
	// if err != nil {
	// 	return err
	// }
	log.Println("=====================================")
	log.Println("Token: ", token)
	log.Println("=====================================")

	// Create the user if not already stored
	user, err := u.CreateUserFromToken(ctx, token)
	if err != nil {
		return err
	}
	err = u.userRepo.CreateUser(ctx, user)
	if err != nil {
		return err
	}

	// Invalidate the verification token
	err = u.VerifyTokenRepo.InvalidateVerifyToken(ctx, token)
	if err != nil {
		return err
	}

	return nil
}

func (u *userUsecase) CreateUserFromToken(ctx context.Context, token string) (*domain.User, error) {
	// Parse the token to get user data
	claims, err := u.authService.ParseToken(token)
	if err != nil {
		return nil, err
	}

	user := &domain.User{
		Firstname: claims.Firstname,
		Lastname:  claims.Lastname,
		Username:  claims.Username,
		Email:     claims.Email,
		Password:  claims.Password, // Assume the password was hashed before being stored in the token
	}

	return user, nil
}
