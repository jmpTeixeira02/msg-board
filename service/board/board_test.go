package board

import (
	"msg-board/protocol"
	"msg-board/service/notifier"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBoard(t *testing.T) {
	{
		t.Logf("Should create new board")
		res := NewBoard("")
		assert.NotNil(t, res)
	}
}

func TestSubscribe(t *testing.T) {
	{
		t.Log("Should subscribe one service")
		userId := ""
		board := MessageBoard{
			Id:            userId,
			Subscriptions: map[string][]protocol.Notifier{},
		}
		n, err := notifier.NewNotifier(protocol.SMS)
		assert.Nil(t, err)
		subscription := protocol.Subscription{
			Subscriber: userId,
			Services:   []protocol.Notifier{n},
		}
		err = board.Subscribe(subscription)
		assert.Len(t, board.Subscriptions[userId], 1)
		assert.Nil(t, err)
	}

	{
		t.Log("Should subscribe multiple services")
		userId := ""
		board := MessageBoard{
			Id:            userId,
			Subscriptions: map[string][]protocol.Notifier{},
		}
		sms, err := notifier.NewNotifier(protocol.SMS)
		assert.Nil(t, err)
		email, err := notifier.NewNotifier(protocol.Email)
		assert.Nil(t, err)
		subscription := protocol.Subscription{
			Subscriber: userId,
			Services:   []protocol.Notifier{sms, email},
		}
		err = board.Subscribe(subscription)
		assert.Nil(t, err)
		assert.Len(t, board.Subscriptions[userId], 2)
	}

	{
		t.Log("Should error on subscribe without notifier")
		userId := ""
		board := MessageBoard{
			Id:            userId,
			Subscriptions: map[string][]protocol.Notifier{},
		}
		subscription := protocol.Subscription{
			Subscriber: userId,
			Services:   nil,
		}
		err := board.Subscribe(subscription)
		assert.Error(t, err)
	}
}

func TestMultipleSubscribe(t *testing.T) {
	{
		t.Log("Should add two subscriptions")
		email, err := notifier.NewNotifier(protocol.Email)
		assert.Nil(t, err)
		board := MessageBoard{
			Id:            "",
			Subscriptions: map[string][]protocol.Notifier{},
		}
		subscription1 := protocol.Subscription{
			Subscriber: "1",
			Services:   []protocol.Notifier{email},
		}
		err = board.Subscribe(subscription1)
		assert.Nil(t, err)
		subscription2 := protocol.Subscription{
			Subscriber: "2",
			Services:   []protocol.Notifier{email},
		}
		err = board.Subscribe(subscription2)
		assert.Nil(t, err)

		assert.Len(t, board.Subscriptions, 2)
	}

	{
		t.Log("Should only add successful subscriptions")
		email, err := notifier.NewNotifier(protocol.Email)
		assert.Nil(t, err)
		board := MessageBoard{
			Id:            "",
			Subscriptions: map[string][]protocol.Notifier{},
		}
		subscription1 := protocol.Subscription{
			Subscriber: "1",
			Services:   []protocol.Notifier{email},
		}
		err = board.Subscribe(subscription1)
		assert.Nil(t, err)
		subscription2 := protocol.Subscription{
			Subscriber: "2",
			Services:   []protocol.Notifier{},
		}
		err = board.Subscribe(subscription2)
		assert.Error(t, err)

		assert.Len(t, board.Subscriptions, 1)
	}
}

func TestUnsubscribe(t *testing.T) {
	{
		t.Log("Should delete on unsubscribe")
		email, err := notifier.NewNotifier(protocol.Email)
		assert.Nil(t, err)
		board := MessageBoard{
			Id:            "",
			Subscriptions: map[string][]protocol.Notifier{},
		}
		subscription := protocol.Subscription{
			Subscriber: "1",
			Services:   []protocol.Notifier{email},
		}
		err = board.Subscribe(subscription)
		assert.Nil(t, err)

		assert.Len(t, board.Subscriptions, 1)
		board.Unsubscribe("1")
		assert.Len(t, board.Subscriptions, 0)
	}
}
