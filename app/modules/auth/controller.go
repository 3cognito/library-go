package auth

import (
	"net/http"

	"github.com/3cognito/library/app/utils"
	"github.com/gin-gonic/gin"
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
