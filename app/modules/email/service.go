package email

import "github.com/3cognito/library/app/config"

func NewEmailService(config config.Config) EmailService {
	return &emailService{
		config: config,
	}
}
