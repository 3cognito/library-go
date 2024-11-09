package auth

import "github.com/3cognito/library/app/modules/email"

func (a *authService) triggerEmailVerificationNotification(recipient, otp string) {
	header := email.EmailTemplates[email.VerifyEmailSubject].Header
	content := email.EmailTemplates[email.VerifyEmailSubject].Content(otp)

	go a.emailService.SendEmailToUser(recipient, header, content)
}
