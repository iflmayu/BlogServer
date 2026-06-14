package email

import (
	"fmt"
	"net/smtp"

	"BlogServer/pkg/config"
)

type Service struct {
	cfg config.Email
}

func NewService(cfg config.Email) *Service {
	return &Service{cfg: cfg}
}

func (s *Service) Send(to string, subject string, body string) error {
	auth := smtp.PlainAuth("", s.cfg.Username, s.cfg.Password, s.cfg.SMTPHost)

	msg := []byte(fmt.Sprintf(
		"To: %s\r\n"+
			"From: %s <%s>\r\n"+
			"Subject: %s\r\n"+
			"Content-Type: text/plain; charset=UTF-8\r\n"+
			"\r\n"+
			"%s\r\n",
		to, s.cfg.FromName, s.cfg.Username, subject, body,
	))

	addr := fmt.Sprintf("%s:%d", s.cfg.SMTPHost, s.cfg.SMTPPort)
	return smtp.SendMail(addr, auth, s.cfg.Username, []string{to}, msg)
}
