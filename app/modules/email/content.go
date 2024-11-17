package email

import "fmt"

//would handle email content better in a more serious app
var EmailTemplates = map[string]EmailTemplate{
	VerifyEmailSubject: {
		Header: VerifyEmailSubject,
		Content: func(params ...interface{}) string {
			return fmt.Sprintf("Hello, please use this otp to verify your email %s", params...)
		},
	},
	ResetPasswordSubject: {
		Header: ResetPasswordSubject,
		Content: func(params ...interface{}) string {
			return fmt.Sprintf("Hello %s, use this otp to reset your password: %s", params...)
		},
	},
	PasswordChanged: {
		Header: PasswordChanged,
		Content: func(params ...interface{}) string {
			return fmt.Sprintf("Hello, your password has been changed at: %s", params...)
		},
	},
}
