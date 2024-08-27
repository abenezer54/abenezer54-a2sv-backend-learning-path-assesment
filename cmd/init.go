package cmd

import (
	"loan-api/delivery/controller/loan_controller"
	"loan-api/delivery/controller/log_controller"
	"loan-api/delivery/controller/user_controller"
	"loan-api/infrastructure/auth"
	"loan-api/infrastructure/bootstrap"
	"loan-api/infrastructure/email"
	"loan-api/repository/loan_repository"
	"loan-api/repository/log_repository"
	"loan-api/repository/refresh_token_repository"
	"loan-api/repository/reset_token_repository"
	"loan-api/repository/user_repository"
	"loan-api/usecase/loan_usecase"
	"loan-api/usecase/log_usecase"
	"loan-api/usecase/user_usecase"
	"time"
)

func InitializeDepenencies() (uc *user_controller.UserController, lc *loan_controller.LoanController, logc log_controller.LogController, env *bootstrap.Env) {
	app := bootstrap.App()
	defer app.CloseDBConnection()
	env = app.Env

	db := app.Mongo.Database(env.DBName)

	userCollection := db.Collection("users")
	refreshTokenCollection := db.Collection("refresh-tokens")
	resetTokenCollection := db.Collection("reset-tokens")

	loanCollection := db.Collection("loans")
	logCollection := db.Collection("logs")

	userRepo := user_repository.NewUserRepository(userCollection)
	refreshTokenRepo := refresh_token_repository.NewRefreshTokenRepository(refreshTokenCollection)
	resetTokenRepo := reset_token_repository.NewResetTokenRepository(resetTokenCollection)

	loanRepo := loan_repository.NewLoanRepository(loanCollection)
	logRepo := log_repository.NewLogRepository(logCollection)

	authService := auth.NewAuthService(refreshTokenRepo, resetTokenRepo, env.AccessTokenSecret, env.RefreshTokenSecret, env.ResetTokenSecret, env.AccessTokenExpiryHour, env.RefreshTokenExpiryHour, env.ResetTokenExpiryHour)

	emailService := email.NewEmailService(env.SMTPServer, env.SMTPPort, env.SMTPUser, env.SMTPPassword, env.FromAddress)

	userUsecase := user_usecase.NewUserUsecase(userRepo, authService, emailService, time.Duration(env.ContextTimeout))
	loanUsecase := loan_usecase.NewLoanUsecase(loanRepo)
	logUsecase := log_usecase.NewLogUsecase(logRepo)

	userController := user_controller.NewUserController(userUsecase, authService, env)
	loanController := loan_controller.NewLoanController(loanUsecase)
	logController := log_controller.NewLogController(logUsecase)

	return userController, loanController, *logController, env

}
