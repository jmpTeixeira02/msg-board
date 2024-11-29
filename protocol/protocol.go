package protocol

type Publisher interface {
	Subscribe(subscription Subscription) error
	Unsubscribe(subscriber string)
	NewMessage(msg string)
}

type Subscriber interface {
	Subscribe(publisher string, services ...NotifyService)
	Unsubscribe(publisher string)
}

type Subscription struct {
	Subscriber string
	Services   []Notifier
}
