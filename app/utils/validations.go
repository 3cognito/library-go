package utils

import (
	"net/http"
	"net/mail"
	"reflect"

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

func NoEmptyFields(param any) bool {
	val := reflect.ValueOf(param)
	if val.Kind() != reflect.Struct {
		return false
	}

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)

		if !field.IsValid() || !field.CanInterface() {
			continue
		}

		if isEmptyValue(field) {
			return false
		}
	}
	return true
}

func isEmptyValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.String:
		return v.Len() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Slice, reflect.Array, reflect.Map, reflect.Chan:
		return v.Len() == 0
	case reflect.Ptr, reflect.Interface:
		return v.IsNil()
	case reflect.Struct:
		for i := 0; i < v.NumField(); i++ {
			if isEmptyValue(v.Field(i)) {
				return true
			}
		}
	}
	return false
}

func IsValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}
