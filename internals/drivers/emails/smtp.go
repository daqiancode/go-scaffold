package emails

import (
	"crypto/tls"
	"errors"
	"net/smtp"
	"strings"

	"github.com/jordan-wright/email"
)

type SMTPMailer struct {
	addr     string
	username string
	password string
	from     string
	authType string
}

func NewSMTPMailer(addr, username, password, from, authType string) *SMTPMailer {
	return &SMTPMailer{
		addr:     addr,
		username: username,
		password: password,
		from:     from,
		authType: authType,
	}
}

type Email struct {
	From    string
	To      []string
	Cc      []string
	Subject string
	Text    string
	HTML    string
}

func (s *SMTPMailer) Send(m Email) error {
	switch s.authType {
	case "starttls":
		return s.SendStartTLS(m)
	case "tls":
		return s.SendTLS(m)
	case "plain":
		return s.SendPlain(m)
	default:
		return s.SendPlain(m)
	}
}
func (s *SMTPMailer) SendPlain(m Email) error {
	e := email.NewEmail()
	e.From = s.from
	e.To = m.To
	e.Cc = m.Cc
	e.Subject = m.Subject
	e.Text = []byte(m.Text)
	e.HTML = []byte(m.HTML)
	if !strings.Contains(s.addr, ":") {
		s.addr += ":25"
	}
	return e.Send(s.addr, smtp.PlainAuth("", s.username, s.password, strings.Split(s.addr, ":")[0]))
}

func (s *SMTPMailer) SendStartTLS(m Email) error {
	e := email.NewEmail()
	e.From = s.from
	e.To = m.To
	e.Cc = m.Cc
	e.Subject = m.Subject
	e.Text = []byte(m.Text)
	e.HTML = []byte(m.HTML)
	if !strings.Contains(s.addr, ":") {
		s.addr += ":587"
	}
	host := strings.Split(s.addr, ":")[0]
	auth := LoginAuth(s.from, s.password)
	tlsconfig := &tls.Config{
		ServerName: host,
	}
	return e.SendWithStartTLS(s.addr, auth, tlsconfig)
}

func (s *SMTPMailer) SendTLS(m Email) error {
	e := email.NewEmail()
	e.From = s.from
	e.To = m.To
	e.Cc = m.Cc
	e.Subject = m.Subject
	e.Text = []byte(m.Text)
	e.HTML = []byte(m.HTML)
	if !strings.Contains(s.addr, ":") {
		s.addr += ":465"
	}
	host := strings.Split(s.addr, ":")[0]
	auth := LoginAuth(s.from, s.password)
	tlsconfig := &tls.Config{
		ServerName: host,
	}
	return e.SendWithTLS(s.addr, auth, tlsconfig)
}

type loginAuth struct {
	username, password string
}

func LoginAuth(username, password string) smtp.Auth {
	return &loginAuth{username, password}
}

func (a *loginAuth) Start(server *smtp.ServerInfo) (string, []byte, error) {
	return "LOGIN", []byte(a.username), nil
}

func (a *loginAuth) Next(fromServer []byte, more bool) ([]byte, error) {
	if more {
		switch string(fromServer) {
		case "Username:":
			return []byte(a.username), nil
		case "Password:":
			return []byte(a.password), nil
		default:
			return nil, errors.New("unknown from server")
		}
	}
	return nil, nil
}
