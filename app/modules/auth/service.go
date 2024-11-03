package auth

import (
	"strings"
	"time"

	"github.com/3cognito/library/app/config"
	"github.com/3cognito/library/app/modules/users"
	"github.com/3cognito/library/app/utils"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func NewAuthService(
	userRepo users.UserRepoInterface,
) AuthServiceInterface {
	return &authService{
		userRepo: userRepo,
	}
}

func (a *authService) SignUp(data SignUpRequest) (LoggedInResponse, error) {
	var res LoggedInResponse
	user := &users.User{
		FirstName:  strings.TrimSpace(data.FirstName),
		MiddleName: strings.TrimSpace(data.MiddleName),
		LastName:   strings.TrimSpace(data.LastName),
		Username:   strings.TrimSpace(data.Username),
		Email:      strings.TrimSpace(data.Email),
		Password:   utils.HashData(data.Password),
		Country:    strings.TrimSpace(data.Country),
		City:       strings.TrimSpace(data.City),
	}

	tx := a.userRepo.BeginTrx()
	if err := a.userRepo.CreateUser(user); err != nil {
		return res, err
	}

	expiryDuration := utils.ParseAccessTokenExpiryTime(config.Configs.AccessTokenExpiryDuration)
	token, tokenErr := generateAccessToken(user.ID, []byte(config.Configs.AppJWTSecret), expiryDuration)
	if tokenErr != nil {
		tx.Rollback()
		return res, tokenErr
	}
	tx.Commit()

	res.Token = token
	utils.ConvertStruct(user, &res.User)

	return res, nil
}

func (a *authService) Login(data LoginRequest) (LoggedInResponse, error) {
	var res LoggedInResponse
	user, err := a.userRepo.GetUserByEmail(data.Email)
	if err != nil || user.IsPasswordCorrect(data.Password) {
		return res, ErrWrongEmailOrPassword
	}

	expiryDuration := utils.ParseAccessTokenExpiryTime(config.Configs.AccessTokenExpiryDuration)
	token, tokenErr := generateAccessToken(user.ID, []byte(config.Configs.AppJWTSecret), expiryDuration)
	if tokenErr != nil {
		return res, tokenErr
	}

	res.Token = token
	utils.ConvertStruct(user, &res.User)

	return res, nil
}

func generateAccessToken(userId uuid.UUID, jwtKey []uint8, expiryTime time.Time) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userId.String(),
		"exp":    expiryTime.Unix(),
	})

	return token.SignedString(jwtKey)
}
