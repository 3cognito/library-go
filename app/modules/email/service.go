package email

import (
	"github.com/3cognito/library/app/config"
	"github.com/resend/resend-go/v2"
)

func NewEmailService(config config.Config) EmailServiceInterface {
	client := resend.NewClient(config.EmailApiKey)
	return &emailService{
		client: *client,
		config: config,
	}
}

func (e *emailService) SendEmailToUser(recipient string, subject string, body string) error {
	if e.config.ENV == config.Dev {
		recipient = "delivered@resend.dev"
	}
	params := &resend.SendEmailRequest{
		From:    e.config.EmailFrom,
		To:      []string{recipient},
		Subject: subject,
		ReplyTo: e.config.ReplyToEmail,
		Text:    body,
	}

	_, err := e.client.Emails.Send(params)
	if err != nil {
		//possibly implement a simple retry mechanism
		return err
	}

	return nil
}
