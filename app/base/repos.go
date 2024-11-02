package base

import "github.com/3cognito/library/app/modules/users"

func (b *base) WithUserRepo() users.UserRepoInterface {
	return users.NewUserRepo(b.db)
}
