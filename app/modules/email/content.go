package email

import "fmt"

//would handle email content better in a more serious app
var EmailTemplates = map[string]EmailTemplate{
	"VERIFY_EMAIL": {
		Header: "Verify your Emaily app",
		Content: func(params ...interface{}) string {
			return fmt.Sprintf("Hello, please verify your email by clicking on this link: %s", params...)
		},
	},
	"RESET_PASSWORD": {
		Header: "Reset Your Password",
		Content: func(params ...interface{}) string {
			return fmt.Sprintf("Hello %s, click this link to reset your password: %s", params...)
		},
	},
}
