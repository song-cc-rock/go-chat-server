package service

import (
	"github.com/jordan-wright/email"
	"github.com/patrickmn/go-cache"
	"go-chat-server/pkg/config"
	"go-chat-server/pkg/utils"
	"net/smtp"
	"time"
)

var (
	verifyCodeCache = cache.New(1*time.Minute, 5*time.Minute)
)

type EmailService interface {
	SendVerifyCode(to string) error
	VerifyCode(mail string, code string) bool
}

type emailService struct {
}

func NewEmailService() EmailService {
	return &emailService{}
}

// SendVerifyCode send verify code to email
func (e *emailService) SendVerifyCode(to string) error {
	code := utils.GenerateVerifyCode()
	err := e.sendVerifyCode(to, code)
	if err != nil {
		return err
	}
	verifyCodeCache.Set(to, code, cache.DefaultExpiration)
	return nil
}

func (e *emailService) VerifyCode(mail string, code string) bool {
	cacheCode, found := verifyCodeCache.Get(mail)
	if !found {
		return false
	}
	if cacheCode != code {
		return false
	}
	return true
}

// sendVerifyCode send verify code to email
func (e *emailService) sendVerifyCode(to string, code string) error {
	em := email.NewEmail()
	fromUser := config.GetString("mail.smtp.username")
	smtpAuthCode := config.GetString("mail.smtp.auth_code")
	smtpHost := config.GetString("mail.smtp.host")
	smtpPort := config.GetString("mail.smtp.port")
	em.From = fromUser
	em.To = []string{to}
	em.Subject = "Go Chat Verify Code"
	em.HTML = []byte(`
		<h1>Verification Code</h1>
		<p>Your verification code is: <strong>` + code + `</strong></p>`)
	err := em.Send(smtpHost+":"+smtpPort, smtp.PlainAuth("", fromUser, smtpAuthCode, smtpHost))
	if err != nil {
		return err
	}
	return nil
}
