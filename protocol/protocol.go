package protocol

type Publisher interface {
	Subscribe(subscriber string, services []Notifier) error
	Unsubscribe(subscriber string)
	NewMessage(msg string)
}

type Subscriber interface {
	Subscribe(publisher string, services ...NotifyService)
	Unsubscribe(publisher string)
}
