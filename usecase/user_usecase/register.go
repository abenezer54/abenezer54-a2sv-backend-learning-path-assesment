package user_usecase

import (
	"context"
	"errors"
	"fmt"
	"loan-api/domain"
	infrastrucutre "loan-api/infrastructure"
	"loan-api/infrastructure/validation"
)

func (u *userUsecase) Register(ctx context.Context, req domain.ResgisterRequest, tokenSecret string, tokenExpiry int) (domain.RegisterResponse, error) {

	err := validation.ValidateEmail(req.Email)
	if err != nil {
		return domain.RegisterResponse{}, err
	}

	existingUser, err := u.userRepo.GetByEmail(ctx, req.Email)
	if err != nil {
		return domain.RegisterResponse{}, err
	}
	if existingUser != nil {
		return domain.RegisterResponse{}, errors.New("email already in use")
	}

	user := &domain.User{
		Firstname: req.Firstname,
		Lastname:  req.Lastname,
		Username:  req.Username,
		Password:  req.Password,
	}
	verifyToken, err := infrastrucutre.GenerateVerificationToken(ctx, *user, tokenSecret, tokenExpiry)
	if err != nil {
		return domain.RegisterResponse{}, err
	}
	verificationURL := fmt.Sprintf("https://localhost:800/verify-email?token=%s", verifyToken)
	err = u.emailService.SendVerificationEmail(ctx, req.Email, verificationURL)
	if err != nil {
		return domain.RegisterResponse{}, err
	}

	u.VerifyTokenRepo.StoreVerifyToken(ctx, verifyToken)

	return domain.RegisterResponse{
		Message: "Verification email sent",
	}, nil

}
