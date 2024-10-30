package utils

import (
	"github.com/gin-gonic/gin"
)

type Response struct {
	Success  bool   `json:"success"`
	Message  string `json:"message"`
	Status   int    `json:"-"`
	Data     any    `json:"data,omitempty"`
	ErrorMsg string `json:"error_msg,omitempty"`
}

func JsonSuccessResponse(ctx *gin.Context, status int, message string, data interface{}) {
	response := Response{Success: true, Message: message, Data: data}
	ctx.JSON(status, response)
}

func JsonErrorResponse(ctx *gin.Context, status int, message, errMsg string) {
	response := Response{Status: status, Success: false, Message: message, ErrorMsg: errMsg}
	ctx.JSON(status, response)
}
