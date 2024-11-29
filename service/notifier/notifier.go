package notifier

import (
	"fmt"
	"msg-board/protocol"
)

func NewNotifiers(services []protocol.NotifyService) ([]protocol.Notifier, error) {
	notifiers := make([]protocol.Notifier, len(services))
	for i := range services {
		n, err := NewNotifier(services[i])
		if err != nil {
			return nil, err
		}
		notifiers[i] = n
	}
	return notifiers, nil
}

func NewNotifier(service protocol.NotifyService) (protocol.Notifier, error) {
	switch protocol.NotifyService(service) {
	case protocol.Email:
		return &EmailNotifier{}, nil
	case protocol.WhatsApp:
		return &WhatsAppNotifier{}, nil
	case protocol.SMS:
		return &SMSNotifier{}, nil
	default:
		return nil, fmt.Errorf("unimplemented notifier")
	}
}

type EmailNotifier struct{}

func (n *EmailNotifier) Send(msg string) {
	fmt.Printf("Email: %s\n", msg)
}

type WhatsAppNotifier struct{}

func (n *WhatsAppNotifier) Send(msg string) {
	fmt.Printf("WhatsApp: %s\n", msg)
}

type SMSNotifier struct{}

func (n *SMSNotifier) Send(msg string) {
	fmt.Printf("SMS: %s\n", msg)
}
