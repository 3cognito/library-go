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
	verifiedEmailRequired := v1.Group("/")
	verifiedEmailRequired.Use(middlewares.VerifiedEmailRequired)

	v1.POST("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

	//authentication and authorization routes
	v1.POST("signup", app.AuthC.SignUp)
	v1.POST("login", app.AuthC.Login)
	v1.PUT("verify-email", middlewares.UserExists, app.AuthC.VerifyEmail)

	//books routes
	books := verifiedEmailRequired.Group("books")
	books.POST("/", app.BooksC.AddBook)
	books.DELETE("/:bookId", app.BooksC.DeleteBook)
	books.GET("/", app.BooksC.GetAuthorBooks)
	books.GET("/:bookId", app.BooksC.GetBook)
	books.PUT("/:bookId/files", app.BooksC.UpdateBookFiles)
	books.PUT("/:bookId/details", app.BooksC.UpdateBookDetails)
}
