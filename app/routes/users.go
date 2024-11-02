package routes

import (
	"github.com/3cognito/library/app/base"
	"github.com/3cognito/library/app/config"
	"github.com/3cognito/library/app/initializers"
	"github.com/gin-gonic/gin"
)

func RouteHandlers(r *gin.Engine) {
	app := base.New(*config.Configs, initializers.DB).LoadControllers()
	v1 := r.Group("api/v1")

	v1.POST("/signup", app.AuthC.SignUp)
}
