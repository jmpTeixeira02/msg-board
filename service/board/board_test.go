package board

import (
	"errors"
	"msg-board/protocol"
	"msg-board/service/notifier"
	"testing"

	"github.com/stretchr/testify/assert"
)

type errorNotifier struct{}

func (n *errorNotifier) Send(msg string) error {
	return errors.New("err")
}

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
		board := NewBoard("")
		sub := protocol.NewSubscription{
			UserId:         userId,
			NotifyServices: []protocol.NotifyService{protocol.SMS},
		}
		err := board.Subscribe(sub)
		assert.Nil(t, err)
		assert.Len(t, board.Subscriptions[userId], 1)
	}

	{
		t.Log("Should subscribe multiple services")
		userId := ""
		board := NewBoard("")
		sub := protocol.NewSubscription{
			UserId:         userId,
			NotifyServices: []protocol.NotifyService{protocol.SMS, protocol.Email},
		}
		err := board.Subscribe(sub)
		assert.Nil(t, err)
		assert.Len(t, board.Subscriptions[userId], 2)
	}

	{
		t.Log("Should error on subscribe with unimplemented notifier")
		userId := ""
		board := NewBoard("")
		sub := protocol.NewSubscription{
			UserId:         userId,
			NotifyServices: []protocol.NotifyService{protocol.NotifyService("")},
		}
		err := board.Subscribe(sub)
		assert.Error(t, err)
	}

	{
		t.Log("Should error on subscribe without notifier")
		userId := ""
		board := NewBoard("")
		sub := protocol.NewSubscription{
			UserId:         userId,
			NotifyServices: []protocol.NotifyService{},
		}
		err := board.Subscribe(sub)
		assert.Error(t, err)
	}
}

func TestMultipleSubscribe(t *testing.T) {
	{
		t.Log("Should add two subscriptions")
		subs := []protocol.NewSubscription{
			{
				UserId:         "1",
				NotifyServices: []protocol.NotifyService{protocol.Email},
			}, {
				UserId:         "2",
				NotifyServices: []protocol.NotifyService{protocol.SMS},
			},
		}
		board := NewBoard("")
		for i := range subs {
			_ = board.Subscribe(subs[i])
		}
		assert.Len(t, board.Subscriptions, 2)
	}

	{
		t.Log("Should only add successful subscriptions")
		subs := []protocol.NewSubscription{
			{
				UserId:         "1",
				NotifyServices: []protocol.NotifyService{protocol.Email},
			}, {
				UserId:         "2",
				NotifyServices: []protocol.NotifyService{},
			},
		}
		board := NewBoard("")
		for i := range subs {
			_ = board.Subscribe(subs[i])
		}
		assert.Len(t, board.Subscriptions, 1)
	}
}

func TestUnsubscribe(t *testing.T) {
	{
		t.Log("Should delete on unsubscribe")
		board := NewBoard("")
		_ = board.Subscribe(protocol.NewSubscription{
			UserId:         "1",
			NotifyServices: []protocol.NotifyService{protocol.Email},
		})
		assert.Len(t, board.Subscriptions, 1)
		board.Unsubscribe("1")
		assert.Len(t, board.Subscriptions, 0)
	}
}

func TestSendMessage(t *testing.T) {
	{
		t.Log("Should message on notifiers")
		notifiers := []protocol.Notifier{&notifier.EmailNotifier{}}
		board := NewBoard("")
		board.Subscriptions["1"] = notifiers
		err := board.SendMessage("")
		assert.Nil(t, err)
	}

	{
		t.Log("Should error on unsuccessful notifiers")
		notifiers := []protocol.Notifier{&errorNotifier{}}
		board := NewBoard("")
		board.Subscriptions["1"] = notifiers
		err := board.SendMessage("")
		assert.Error(t, err)
	}
}
