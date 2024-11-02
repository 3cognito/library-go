package auth

import (
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

	res, err := a.authService.SignUp(params)
	if err != nil {
		utils.JsonErrorResponse(ctx, 400, "signup unsuccessful", err.Error())
		return
	}

	utils.JsonSuccessResponse(ctx, 201, "signup successful", res)
}
