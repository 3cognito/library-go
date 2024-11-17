package email

import "fmt"

//would handle email content better in a more serious app
var EmailTemplates = map[string]EmailTemplate{
	VerifyEmailSubject: {
		Header: "Verify your Email",
		Content: func(params ...interface{}) string {
			return fmt.Sprintf("Hello, please use this otp to verify your email %s", params...)
		},
	},
	"RESET_PASSWORD": {
		Header: "Reset Your Password",
		Content: func(params ...interface{}) string {
			return fmt.Sprintf("Hello %s, use this otp to reset your password: %s", params...)
		},
	},
}
