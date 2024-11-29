package board

import (
	"errors"
	"fmt"
	"msg-board/protocol"
	"strings"
)

type MessageBoard struct {
	Id            string
	Subscriptions map[string][]protocol.Notifier // userId -> []notifiers
	Private       bool
	Password      string
}

func isPrivate(pw string) bool {
	return strings.TrimSpace(pw) != ""
}

func NewBoard(id string, pw string) protocol.Publisher {
	return &MessageBoard{
		Id:            id,
		Subscriptions: map[string][]protocol.Notifier{},
		Private:       isPrivate(pw),
		Password:      pw,
	}
}

func (b *MessageBoard) Subscribe(subscription protocol.Subscription, pw string) error {
	if b.Private && pw != b.Password {
		return errors.New("invalid password")
	}
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
