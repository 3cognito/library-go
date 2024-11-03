package base

import (
	"github.com/3cognito/library/app/modules/auth"
	"github.com/3cognito/library/app/modules/email"
)

func (b *base) WithAuthService() auth.AuthServiceInterface {
	return auth.NewAuthService(b.WithUserRepo())
}

func (b *base) WithEmailService() email.EmailService {
	return email.NewEmailService(b.configs)
}
