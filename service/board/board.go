package board

import (
	"errors"
	"fmt"
	"msg-board/protocol"
)

type MessageBoard struct {
	Id            string
	Subscriptions map[string][]protocol.Notifier // userId -> map[Notifier] msgCh
}

func NewBoard(id string) protocol.Publisher {
	return &MessageBoard{
		Id:            id,
		Subscriptions: map[string][]protocol.Notifier{},
	}
}

func (b *MessageBoard) Subscribe(subscription protocol.Subscription) error {
	if len(subscription.Services) < 1 {
		return errors.New("subscription must have notifiers")
	}
	b.Subscriptions[subscription.Subscriber] = subscription.Services
	return nil
}

func (b *MessageBoard) Unsubscribe(user string) {
	delete(b.Subscriptions, user)
}

func (b *MessageBoard) NewMessage(msg string) {
	for user, sub := range b.Subscriptions {
		for _, n := range sub {
			n.Send(fmt.Sprintf("User %s %s", user, msg))
		}
	}
}
