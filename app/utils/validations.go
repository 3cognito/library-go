package utils

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func ValidParams(ctx *gin.Context, params interface{}) bool {
	if err := ctx.ShouldBindBodyWith(&params, binding.JSON); err != nil {
		JsonErrorResponse(ctx, 400, GENERIC_ERROR, INVALID_INPUT)
		return false
	}

	ConvertDataToMap(params)

	return true
}
