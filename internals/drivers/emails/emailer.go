package emails

import (
	"fmt"
	"strings"

	"github.com/daqiancode/env"
	"github.com/flosch/pongo2/v4"
)

type Emailer struct {
	protocol string
	host     string
	username string
	password string
	authType string
}

func NewEmailerEnv() *Emailer {
	return &Emailer{
		protocol: env.Get("EMAIL_PROTOCOL"),
		host:     env.Get("EMAIL_HOST"),
		username: env.Get("EMAIL_USERNAME"),
		password: env.Get("EMAIL_PASSWORD"),
		authType: strings.ToLower(env.Get("EMAIL_AUTH_TYPE")),
	}
}

func (s *Emailer) Send(m Email) error {
	fmt.Printf("send email from %v\n", *s)
	// send email with smtp
	emailer := NewSMTPMailer(s.host, s.username, s.password, s.username, s.authType)
	return emailer.Send(m)
}

const (
	CodeTypeSignupEmailCode = 1
)

func (s *Emailer) SendCode(subject, topic, to, code, name string, ttlMinutes int) error {
	tpl, err := pongo2.FromFile("emails/code.html")
	if err != nil {
		return err
	}
	preheader := "Your verification code is " + code + "."
	if name == "" {
		name = "customer"
	}
	data := map[string]any{
		"topic":        topic,
		"ttl":          ttlMinutes,
		"code":         code,
		"name":         name,
		"title":        "Verfication code from " + env.Get("COMPANY_NAME"),
		"preheader":    preheader,
		"COMPANY_NAME": env.Get("COMPANY_NAME"),
		"COMPANY_URL":  env.Get("COMPANY_URL"),
	}
	html, err := tpl.Execute(data)
	if err != nil {
		return err
	}
	return s.Send(Email{
		To:      []string{to},
		From:    s.username,
		Subject: subject,
		HTML:    html,
		Text:    preheader,
	})
}
