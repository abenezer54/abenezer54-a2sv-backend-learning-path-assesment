package user_controller

import (
	"loan-api/domain"
	"loan-api/infrastructure/bootstrap"
)

type UserController struct {
	userUsecase domain.UserUsecase
	authService domain.AuthService
	Env         *bootstrap.Env
}

func NewUserController(userUsecase domain.UserUsecase, authService domain.AuthService, env *bootstrap.Env) *UserController {
	return &UserController{
		userUsecase: userUsecase,
		authService: authService,
		Env:         env,
	}
}
