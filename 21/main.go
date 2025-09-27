package main

import "fmt"

type Notifier interface {
	NotifyMessage(text string) error
}

type GmailSender struct {
}

func (s *GmailSender) SendMsg(mail, text string) error {
	fmt.Printf("Msg was sent on `%s` mail\n", mail)

	return nil
}

type GmailSenderAdapter struct {
	gs   *GmailSender
	mail string
}

func NewGmailSenderAdapter(gs *GmailSender, mail string) *GmailSenderAdapter {
	return &GmailSenderAdapter{
		gs:   gs,
		mail: mail,
	}
}

func (a *GmailSenderAdapter) NotifyMessage(text string) error {
	return a.gs.SendMsg(a.mail, text)
}

func main() {
	gmailSender := &GmailSender{}
	adapter := NewGmailSenderAdapter(gmailSender, "user@example.com")

	// Теперь можно использовать GmailSender через интерфейс Notifier
	var notifier Notifier = adapter
	_ = notifier.NotifyMessage("Hello, World!")
}
