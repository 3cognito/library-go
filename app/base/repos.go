package base

import (
	"github.com/3cognito/library/app/modules/books"
	"github.com/3cognito/library/app/modules/otp"
	"github.com/3cognito/library/app/modules/users"
)

func (b *base) WithUserRepo() users.UserRepoInterface {
	return users.NewUserRepo(b.db)
}

func (b *base) WithBookRepo() books.BookRepoInterface {
	return books.NewBookRepo(b.db)
}

func (b *base) WithOtpRepo() otp.OtpRepoInterface {
	return otp.NewOtpRepo(b.db)
}
