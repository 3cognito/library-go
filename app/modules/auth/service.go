package auth

import (
	"strings"
	"time"

	"github.com/3cognito/library/app/config"
	"github.com/3cognito/library/app/modules/email"
	"github.com/3cognito/library/app/modules/otp"
	"github.com/3cognito/library/app/modules/users"
	"github.com/3cognito/library/app/utils"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

func NewAuthService(
	userRepo users.UserRepoInterface,
	otpService otp.OtpServiceInterface,
	emailService email.EmailServiceInterface,
) AuthServiceInterface {
	return &authService{
		userRepo:     userRepo,
		otpService:   otpService,
		emailService: emailService,
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

	expiryDuration := utils.ParseAccessTokenExpiryDuration(config.Configs.AccessTokenExpiryDuration)
	token, tokenErr := generateAccessToken(user.ID, []byte(config.Configs.AppJWTSecret), expiryDuration)
	if tokenErr != nil {
		tx.Rollback()
		return res, tokenErr
	}

	res.Token = token
	utils.ConvertStruct(user, &res.User)

	otp, err := a.otpService.CreateOtp(user.ID, otp.EmailVerifcation, ADayHence)
	if err != nil {
		tx.Rollback()
	}
	tx.Commit()

	a.triggerEmailVerificationNotification(user.Email, otp.Value)

	return res, nil
}

func (a *authService) Login(data LoginRequest) (LoggedInResponse, error) {
	var res LoggedInResponse
	user, err := a.userRepo.GetUserByEmail(strings.TrimSpace(data.Email))
	if err != nil || !user.IsPasswordCorrect(data.Password) {
		return res, ErrWrongEmailOrPassword
	}

	expiryDuration := utils.ParseAccessTokenExpiryDuration(config.Configs.AccessTokenExpiryDuration)
	token, tokenErr := generateAccessToken(user.ID, []byte(config.Configs.AppJWTSecret), expiryDuration)
	if tokenErr != nil {
		return res, tokenErr
	}

	res.Token = token
	utils.ConvertStruct(user, &res.User)

	return res, nil
}

func (a *authService) VerifyEmail(data VerifyEmailRequest) error {
	user, err := a.userRepo.GetUserByID(data.UserID)
	if err != nil {
		return ErrAccountNotFound
	}

	otp, err := a.otpService.GetOtpByUseCase(data.UserID, otp.EmailVerifcation)
	if err != nil || otp.Value != data.Otp {
		return ErrOtpExpiredOrInvalid
	}

	if err := a.otpService.InValidateOtp(otp); err != nil {
		//log error - returning user readable error
		return ErrOtpExpiredOrInvalid
	}

	now := time.Now()
	user.EmailVerifiedAt = &now
	if err := a.userRepo.UpdateUser(user); err != nil {
		return err
	}

	return nil
}

func (a *authService) ForgotPassword(email string) error {
	user, userErr := a.userRepo.GetUserByEmail(email)
	if userErr != nil {
		return ErrAccountNotFound
	}

	otp, err := a.otpService.CreateOtp(user.ID, otp.PasswordReset, ADayHence)
	if err != nil {
		//log issue
		return err
	}

	a.triggerPasswordResetNotification(user.Email, otp.Value)

	return nil

}

func generateAccessToken(userId uuid.UUID, jwtKey []uint8, expiryDuration time.Duration) (string, error) {
	claims := jwt.RegisteredClaims{
		Subject:   userId.String(),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiryDuration)),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString(jwtKey)
}
