package auth

import (
	"github.com/3cognito/library/app/modules/users"
	"github.com/gin-gonic/gin"
)

type authService struct {
	userRepo users.UserRepoInterface
}

type AuthServiceInterface interface {
	SignUp(data SignUpRequest) (LoggedInResponse, error)
	Login(data LoginRequest) (LoggedInResponse, error)
}

type authController struct {
	authService AuthServiceInterface
}

type AuthControllerInterface interface {
	SignUp(ctx *gin.Context)
	Login(ctx *gin.Context)
}
