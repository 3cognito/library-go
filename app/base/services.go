package base

import "github.com/3cognito/library/app/modules/auth"

func (b *base) WithAuthService() auth.AuthServiceInterface {
	return auth.NewAuthService(b.WithUserRepo())
}
