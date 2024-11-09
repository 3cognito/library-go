package middlewares

import (
	"net/http"
	"strings"

	"github.com/3cognito/library/app/config"
	"github.com/3cognito/library/app/initializers"
	"github.com/3cognito/library/app/modules/users"
	"github.com/3cognito/library/app/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func ValidateAuthToken(ctx *gin.Context) (*users.User, error) {
	auth := ctx.GetHeader("Authorization")
	extractedToken, extractionErr := extractBearerToken(auth)
	if extractionErr != nil {
		return nil, extractionErr
	}

	claims := jwt.RegisteredClaims{}
	token, err := jwt.ParseWithClaims(extractedToken, &claims, func(token *jwt.Token) (any, error) {
		return []byte(config.Configs.AppJWTSecret), nil
	})

	if err != nil || !token.Valid {
		return nil, ErrInvalidOrExpiredToken
	}

	userId := uuid.MustParse(claims.Subject)

	user, userErr := users.NewUserRepo(initializers.DB).GetUserByID(userId)
	if userErr != nil || user.DeletedAt.Valid {
		return nil, ErrAccountNotFoundOrDeleted
	}

	return user, nil
}

func UserExists(ctx *gin.Context) {
	user, validationErr := ValidateAuthToken(ctx)
	if validationErr != nil {
		utils.JsonErrorResponse(ctx, http.StatusUnauthorized, "unauthorized", validationErr.Error())
	}

	ctx.Set("userId", user.ID.String())
	ctx.Next()
}

func VerifiedEmailRequired(ctx *gin.Context) {
	user, validationErr := ValidateAuthToken(ctx)
	if validationErr != nil {
		utils.JsonErrorResponse(ctx, http.StatusUnauthorized, "unauthorized", ErrAccountNotFoundOrDeleted.Error())
		ctx.Abort()
		return
	}
	if user.EmailVerifiedAt == nil {
		utils.JsonErrorResponse(ctx, http.StatusUnauthorized, "unauthorized", ErrEmailNotVerified.Error())
		ctx.Abort()
		return
	}

	ctx.Set("userId", user.ID.String())
	ctx.Next()
}

func extractBearerToken(header string) (string, error) {
	if header == "" {
		return "", ErrAuthRequired
	}

	jwtToken := strings.Split(header, " ")
	if len(jwtToken) != 2 {
		return "", ErrInvalidOrExpiredToken
	}

	return jwtToken[1], nil
}
