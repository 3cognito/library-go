package base

import (
	"github.com/3cognito/library/app/modules/auth"
	"github.com/3cognito/library/app/modules/email"
	"github.com/3cognito/library/app/modules/otp"
)

func (b *base) WithAuthService() auth.AuthServiceInterface {
	return auth.NewAuthService(b.WithUserRepo(), b.WithOtpService())
}

func (b *base) WithEmailService() email.EmailService {
	return email.NewEmailService(b.configs)
}

func (b *base) WithOtpService() otp.OtpServiceInterface {
	return otp.NewOtpService(b.WithOtpRepo())
}
