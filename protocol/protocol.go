package protocol

type Publisher interface {
	Subscribe(subscription NewSubscription) error
	Unsubscribe(subscriber string)
	SendMessage(msg string)
}

type Subscriber interface {
	Subscribe(publisher string, services ...NotifyService)
	Unsubscribe(publisher string)
}

type NewSubscription struct {
	UserId         string
	NotifyServices []NotifyService
}

type Subscription struct {
	NotifyServices []Notifier
	MsgCh          chan string
}
