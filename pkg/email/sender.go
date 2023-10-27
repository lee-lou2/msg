package email

import (
	"github.com/lee-lou2/msg/models"
	"gopkg.in/gomail.v2"
	"strconv"
)

// SendEmail 이메일 개별 전송
func SendEmail(users []*models.User, subject, message string, configs ...Configs) error {
	// 이메일 발송
	cfg, err := setConfigs(configs...)
	if err != nil {
		return err
	}

	for _, u := range users {
		email := u.Email
		m := gomail.NewMessage()
		m.SetHeader("From", cfg.User)
		m.SetHeader("To", email)
		m.SetHeader("Subject", subject)
		m.SetBody("text/html", message)
		portStr := cfg.Port
		port, err := strconv.Atoi(portStr)
		if err != nil {
			return err
		}
		d := gomail.NewDialer(cfg.Host, port, cfg.User, cfg.Password)
		if err := d.DialAndSend(m); err != nil {
			return err
		}
	}
	return nil
}
