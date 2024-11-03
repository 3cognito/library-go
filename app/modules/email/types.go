package email

import (
	"github.com/3cognito/library/app/config"
	"github.com/resend/resend-go/v2"
)

type emailService struct {
	client resend.Client
	config config.Config
}

type EmailService interface {
	SendEmailToUser(recipient string, subject string, body string) error
}

type EmailTemplate struct {
	Header  string
	Content func(params ...interface{}) string
}
