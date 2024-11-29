package board

import (
	"errors"
	"msg-board/protocol"
	"msg-board/service/notifier"
)

type MessageBoard struct {
	Id            string
	MsgCh         chan string
	Subscriptions map[string][]protocol.Notifier // userId -> NotifiyServices
}

func NewBoard(id string) MessageBoard {
	return MessageBoard{
		Id:            id,
		MsgCh:         make(chan string, 1),
		Subscriptions: map[string][]protocol.Notifier{},
	}
}

func (b *MessageBoard) Subscribe(sub protocol.NewSubscription) error {
	if len(sub.NotifyServices) < 1 {
		return errors.New("a subscription must have notifiers")
	}
	notifiers := make([]protocol.Notifier, len(sub.NotifyServices))
	for i := range sub.NotifyServices {
		notifier, err := notifier.NewNotifier(sub.NotifyServices[i])
		if err != nil {
			return err
		}
		notifiers[i] = notifier
	}
	b.Subscriptions[sub.UserId] = notifiers
	return nil
}

func (b *MessageBoard) Unsubscribe(user string) {
	delete(b.Subscriptions, user)
}

func (b *MessageBoard) SendMessage(msg string) error {
	var err error
	for _, notifiers := range b.Subscriptions {
		for _, notifier := range notifiers {
			iErr := notifier.Send(msg)
			if iErr != nil {
				err = errors.Join(err, iErr)
			}
		}
	}
	return err
}
