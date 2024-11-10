package base

import (
	"github.com/3cognito/library/app/modules/auth"
	"github.com/3cognito/library/app/modules/books"
	"github.com/3cognito/library/app/modules/cloudinary"
	"github.com/3cognito/library/app/modules/email"
	"github.com/3cognito/library/app/modules/otp"
)

func (b *base) WithAuthService() auth.AuthServiceInterface {
	return auth.NewAuthService(b.WithUserRepo(), b.WithOtpService(), b.WithEmailService())
}

func (b *base) WithEmailService() email.EmailServiceInterface {
	return email.NewEmailService(b.configs)
}

func (b *base) WithOtpService() otp.OtpServiceInterface {
	return otp.NewOtpService(b.WithOtpRepo())
}

func (b *base) WithCloudinaryService() cloudinary.CloudinaryServiceInterface {
	return cloudinary.NewService(b.configs)
}

func (b *base) WithBookService() books.ServiceInterface {
	return books.NewService(b.WithBookRepo(), b.WithCloudinaryService())
}
