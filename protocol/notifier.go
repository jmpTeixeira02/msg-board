package protocol

type NotifyService string

const (
	Email    NotifyService = "email"
	WhatsApp NotifyService = "whatsapp"
	SMS      NotifyService = "sms"
)

type Notifier interface {
	SendNotification(msg string)
}