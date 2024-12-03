package protocol

type Publisher interface {
	Subscribe(subscription Subscription, pw string) error
	Unsubscribe(publisher string, subscriber string)
	NewMessage(msg string)
}

type Subscriber interface {
	Subscribe(publisher string, services ...NotifyService)
	Unsubscribe(publisher string)
}

type Subscribing struct {
	User     string
	Services []NotifyService
}

type Subscription struct {
	Subscriber Subscribing
	Publisher  string
}
