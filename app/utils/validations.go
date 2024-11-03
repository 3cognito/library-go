package utils

import (
	"net/http"
	"net/mail"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func ValidParams(ctx *gin.Context, params interface{}) bool {
	if err := ctx.ShouldBindBodyWith(&params, binding.JSON); err != nil {
		JsonErrorResponse(ctx, http.StatusUnprocessableEntity, GENERIC_ERROR, INVALID_INPUT)
		return false
	}

	ConvertDataToMap(params)

	return true
}

func IsValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
