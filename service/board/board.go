package board

import (
	"errors"
	"fmt"
	"msg-board/protocol"
)

type MessageBoard struct {
	Id            string
	MsgCh         chan string
	Subscriptions map[string][]protocol.Notifier // userId -> map[Notifier] msgCh
}

func NewBoard(id string) protocol.Publisher {
	return &MessageBoard{
		Id:            id,
		MsgCh:         make(chan string, 1),
		Subscriptions: map[string][]protocol.Notifier{},
	}
}

func (b *MessageBoard) Subscribe(userId string, notifiers []protocol.Notifier) error {
	if len(notifiers) < 1 {
		return errors.New("subscription must have notifiers")
	}
	b.Subscriptions[userId] = notifiers
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
