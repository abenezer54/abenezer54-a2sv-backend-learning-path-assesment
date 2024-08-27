package router

import (
	"loan-api/delivery/controller/loan_controller"
	"loan-api/delivery/controller/log_controller"
	"loan-api/delivery/controller/user_controller"
	"loan-api/infrastructure/auth"
	"loan-api/infrastructure/bootstrap"

	"github.com/gin-gonic/gin"
)

func SetRouter(router *gin.Engine, uc *user_controller.UserController, lc *loan_controller.LoanController, logc log_controller.LogController, env *bootstrap.Env) {
	// User routes
	user_routes := router.Group("/users")
	user_routes.POST("/register", uc.Register)
	user_routes.GET("/verify-email", uc.VerifyEmail)
	user_routes.POST("/login", uc.Login)
	user_routes.POST("/token/refresh", auth.JwtAuthMiddleware(env.AccessTokenSecret), uc.RefreshTokens)
	user_routes.POST("/profile", auth.JwtAuthMiddleware(env.AccessTokenSecret), uc.GetProfile)
	user_routes.POST("/forgot-password", uc.ForgotPassword)
	router.POST("/reset-password", uc.ResetPassword)
	admin_routes := router.Group("/admin")
	admin_routes.Use(auth.AdminRoleMiddleware())
	{
		// users
		admin_routes.GET("/users", uc.ViewAllUsers)
		admin_routes.DELETE("/users/:id", uc.DeleteUserAccount)

		// loans
		admin_routes.GET("/loans", lc.GetAllLoans)
		admin_routes.PATCH("/loans/:id/status", lc.UpdateLoanStatus)
		admin_routes.DELETE("/loans/:id", lc.DeleteLoan)

		// logs
		admin_routes.GET("/logs", logc.GetLogs)
	}

	//Loan routes
	router.POST("/loans", auth.JwtAuthMiddleware(env.AccessTokenSecret), lc.ApplyForLoan)
	router.GET("/loans/:id", auth.JwtAuthMiddleware(env.AccessTokenSecret), lc.GetLoanStatus)

}
