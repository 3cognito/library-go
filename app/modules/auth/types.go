package auth

import (
	"github.com/3cognito/library/app/modules/email"
	"github.com/3cognito/library/app/modules/otp"
	"github.com/3cognito/library/app/modules/users"
	"github.com/gin-gonic/gin"
)

type authService struct {
	userRepo     users.UserRepoInterface
	otpService   otp.OtpServiceInterface
	emailService email.EmailServiceInterface
}

type AuthServiceInterface interface {
	SignUp(data SignUpRequest) (LoggedInResponse, error)
	Login(data LoginRequest) (LoggedInResponse, error)
	VerifyEmail(data VerifyEmailRequest) error
	ForgotPassword(email string) error
	ResetPassword(data ResetPasswordRequest) (LoggedInResponse, error)
}

type authController struct {
	authService AuthServiceInterface
}

type AuthControllerInterface interface {
	SignUp(ctx *gin.Context)
	Login(ctx *gin.Context)
	VerifyEmail(ctx *gin.Context)
	ForgotPassword(ctx *gin.Context)
	ResetPassword(ctx *gin.Context)
}
