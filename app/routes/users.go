package routes

import (
	"github.com/3cognito/library/app/base"
	"github.com/3cognito/library/app/config"
	"github.com/3cognito/library/app/initializers"
	"github.com/3cognito/library/app/middlewares"
	"github.com/gin-gonic/gin"
)

func RouteHandlers(r *gin.Engine) {
	app := base.New(*config.Configs, initializers.DB).LoadControllers()
	v1 := r.Group("api/v1")
	authRequired := v1.Group("/")
	authRequired.Use(middlewares.AuthMiddleware)

	v1.POST("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

	//authentication and authorization routes
	v1.POST("signup", app.AuthC.SignUp)
	v1.POST("login", app.AuthC.Login)
	authRequired.PUT("verify-email", app.AuthC.VerifyEmail)
}
