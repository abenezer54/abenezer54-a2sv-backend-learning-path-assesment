package user_usecase

import (
	"time"

	"loan-api/domain"
)

type userUsecase struct {
	userRepo        domain.UserRepository
	VerifyTokenRepo domain.VerifyTokenRepository
	authService     domain.AuthService
	emailService    domain.EmailService
	contextTimeout  time.Duration
}

func NewUserUsecase(userRepository domain.UserRepository, verifyTokenRepo domain.VerifyTokenRepository, authService domain.AuthService, emailService domain.EmailService, timeout time.Duration) *userUsecase {
	return &userUsecase{
		userRepo:       userRepository,
		VerifyTokenRepo: verifyTokenRepo,
		emailService:   emailService,
		authService:    authService,
		contextTimeout: timeout,
	}
}
