package service

import (
	"fmt"
	"msg-board/protocol"
)

func NewNotifier(service protocol.NotifyService) (protocol.Notifier, error) {
	switch protocol.NotifyService(service) {
	case protocol.Email:
		return &EmailNotifier{}, nil
	case protocol.WhatsApp:
		return &WhatsAppNotifier{}, nil
	case protocol.SMS:
		return &SMSNotifier{}, nil
	default:
		return nil, fmt.Errorf("Unimplemented notifier")
	}
}

type EmailNotifier struct{}

func (n *EmailNotifier) SendNotification(msg string) {
	fmt.Printf("New Email: %s\n", msg)
}

type WhatsAppNotifier struct{}

func (n *WhatsAppNotifier) SendNotification(msg string) {
	fmt.Printf("New Message via WhatsApp: %s\n", msg)
}

type SMSNotifier struct{}

func (n *SMSNotifier) SendNotification(msg string) {
	fmt.Printf("New Message via SMS: %s\n", msg)
}
