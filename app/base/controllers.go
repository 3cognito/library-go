package base

import "github.com/3cognito/library/app/modules/auth"

func (b *base) WithAuthController() auth.AuthControllerInterface {
	return auth.NewAuthController(b.WithAuthService())
}
