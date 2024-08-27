package router

import (
	"loan-api/delivery/controller/user_controller"
	"loan-api/infrastructure/bootstrap"

	"github.com/gin-gonic/gin"
)

func SetRouter(router *gin.Engine, uc *user_controller.UserController, env *bootstrap.Env) {
	// User routes
	user_routes := router.Group("/users")
	user_routes.POST("/register", uc.Register)
		
	// router.GET("/verify-email", uc.VerifyEmail)
	// router.POST("/login", uc.Login)
	// router.POST("/refresh", uc.RefreshTokens)
	// router.GET("/logout", auth.JwtAuthMiddleware(env.AccessTokenSecret), uc.Logout)
	// router.POST("/forgot-password", uc.ForgotPassword)
	// router.POST("/reset-password", uc.ResetPassword)
	// router.POST("/generate", auth.JwtAuthMiddleware(env.AccessTokenSecret), bc.GenerateContent)
	// router.PUT("/updateUser", auth.JwtAuthMiddleware(env.AccessTokenSecret), uc.UpdateUser)
	// router.PATCH("/user/promote-demote", auth.JwtAuthMiddleware(env.AccessTokenSecret), uc.PromoteDemote)

	// // Blog routes
	// r.Use(auth.JwtAuthMiddleware(env.AccessTokenSecret))
	// {
	// 	r.POST("/create", bc.CreateBlog)
	// 	r.GET("/", bc.GetBlogs)
	// 	r.GET("/:id", bc.GetBlogByID)
	// 	r.PUT("/:id", bc.UpdateBlog)
	// 	r.DELETE("/:id", bc.DeleteBlog)
	// 	r.GET("/search", bc.SearchBlog)
	// 	r.POST("/filters", bc.FilterBlog)

	// }
}
