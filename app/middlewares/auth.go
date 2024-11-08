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

func AuthMiddleware(ctx *gin.Context) {
	auth := ctx.GetHeader("Authorization")
	extractedToken, extractionErr := extractBearerToken(auth)
	if extractionErr != nil {
		utils.JsonErrorResponse(ctx, http.StatusUnauthorized, "unauthorized", extractionErr.Error())
		ctx.Abort()
		return
	}

	claims := jwt.RegisteredClaims{}
	token, err := jwt.ParseWithClaims(extractedToken, &claims, func(token *jwt.Token) (any, error) {
		return config.Configs.AppJWTSecret, nil
	})

	if err != nil || !token.Valid {
		utils.JsonErrorResponse(ctx, http.StatusUnauthorized, "unauthorized", ErrInvalidToken.Error())
		ctx.Abort()
		return
	}

	userId := uuid.MustParse(claims.Subject)

	if verifyErr := verifyAuthUser(userId); verifyErr != nil {
		utils.JsonErrorResponse(ctx, http.StatusUnauthorized, "unauthorized", verifyErr.Error())
		ctx.Abort()
		return
	}

	ctx.Set("userId", userId)
	ctx.Next()
}

func verifyAuthUser(userId uuid.UUID) error {
	user, userErr := users.NewUserRepo(initializers.DB).GetUserByID(userId)
	if userErr != nil || user.DeletedAt.Valid {
		return ErrAccountNotFoundOrDeleted
	}

	if user.EmailVerifiedAt == nil {
		return ErrEmailNotVerified
	}

	return nil
}

func extractBearerToken(header string) (string, error) {
	if header == "" {
		return "", ErrAuthRequired
	}

	jwtToken := strings.Split(header, " ")
	if len(jwtToken) != 2 {
		return "", ErrInvalidToken
	}

	return jwtToken[1], nil
}
