package auth

import (
	"net/http"

	"github.com/3cognito/library/app/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func NewAuthController(
	authService AuthServiceInterface,
) AuthControllerInterface {
	return &authController{
		authService: authService,
	}
}

func (a *authController) SignUp(ctx *gin.Context) {
	var params SignUpRequest

	if !utils.ValidParams(ctx, &params) {
		return
	}

	if !utils.IsValidEmail(params.Email) {
		utils.JsonErrorResponse(ctx, http.StatusBadRequest, "signup unsuccessful", ErrInvalidEmail.Error())
		return
	}

	res, err := a.authService.SignUp(params)
	if err != nil {
		utils.JsonErrorResponse(ctx, http.StatusBadRequest, "signup unsuccessful", err.Error())
		return
	}

	utils.JsonSuccessResponse(ctx, http.StatusCreated, "signup successful", res)
}

func (a *authController) Login(ctx *gin.Context) {
	var params LoginRequest

	if !utils.ValidParams(ctx, &params) {
		return
	}

	if !utils.IsValidEmail(params.Email) {
		utils.JsonErrorResponse(ctx, http.StatusBadRequest, "login unsuccessful", ErrInvalidEmail.Error())
		return
	}

	res, err := a.authService.Login(params)
	if err != nil {
		utils.JsonErrorResponse(ctx, http.StatusBadRequest, "login unsuccessful", err.Error())
		return
	}

	utils.JsonSuccessResponse(ctx, http.StatusOK, "login successful", res)
}

func (a *authController) VerifyEmail(ctx *gin.Context) {
	parsedUserId, parseErr := uuid.Parse(ctx.GetString("userId"))
	if parseErr != nil {
		utils.JsonErrorResponse(ctx, http.StatusBadRequest, "verification unsuccessful", parseErr.Error())
		return
	}

	var params VerifyEmailRequest
	params.UserID = parsedUserId

	if !utils.ValidParams(ctx, &params) {
		return
	}

	err := a.authService.VerifyEmail(params)
	if err != nil {
		utils.JsonErrorResponse(ctx, http.StatusBadRequest, "verification unsuccessful", err.Error())
		return
	}

	utils.JsonSuccessResponse(ctx, http.StatusOK, "verification successful", nil)
}

func (a *authController) ForgotPassword(ctx *gin.Context) {
	var params ForgotPasswordRequest

	if !utils.ValidParams(ctx, &params) {
		return
	}

	if !utils.IsValidEmail(params.Email) {
		utils.JsonErrorResponse(ctx, http.StatusBadRequest, "invalid email address", ErrInvalidEmail.Error())
		return
	}

	err := a.authService.ForgotPassword(params.Email)
	if err != nil && err != ErrAccountNotFound {
		utils.JsonErrorResponse(ctx, http.StatusBadRequest, "error sending password reset email", err.Error())
		return
	}

	utils.JsonSuccessResponse(ctx, http.StatusOK, "password reset otp sent to provided email address", nil)
}

func (a *authController) ResetPassword(ctx *gin.Context) {
	var params ResetPasswordRequest

	if !utils.ValidParams(ctx, &params) {
		return
	}

	if !utils.IsValidEmail(params.Email) {
		utils.JsonErrorResponse(ctx, http.StatusBadRequest, "invalid email address", ErrInvalidEmail.Error())
		return
	}

	res, err := a.authService.ResetPassword(params)
	if err != nil {
		utils.JsonErrorResponse(ctx, http.StatusBadRequest, "reset password unsuccessful", err.Error())
		return
	}

	utils.JsonSuccessResponse(ctx, http.StatusOK, "password reset successful", res)
}
