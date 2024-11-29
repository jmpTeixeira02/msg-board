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
			MsgCh:         make(chan string, 1),
			Subscriptions: map[string][]protocol.Notifier{},
		}
		n, err := notifier.NewNotifier(protocol.SMS)
		assert.Nil(t, err)
		err = board.Subscribe(userId, []protocol.Notifier{n})
		assert.Nil(t, err)
		assert.Len(t, board.Subscriptions[userId], 1)
	}

	{
		t.Log("Should subscribe multiple services")
		userId := ""
		board := MessageBoard{
			Id:            userId,
			MsgCh:         make(chan string, 1),
			Subscriptions: map[string][]protocol.Notifier{},
		}
		sms, err := notifier.NewNotifier(protocol.SMS)
		assert.Nil(t, err)
		email, err := notifier.NewNotifier(protocol.Email)
		assert.Nil(t, err)
		err = board.Subscribe(userId, []protocol.Notifier{sms, email})
		assert.Nil(t, err)
		assert.Len(t, board.Subscriptions[userId], 2)
	}

	{
		t.Log("Should error on subscribe without notifier")
		userId := ""
		board := MessageBoard{
			Id:            userId,
			MsgCh:         make(chan string, 1),
			Subscriptions: map[string][]protocol.Notifier{},
		}
		err := board.Subscribe(userId, nil)
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
			MsgCh:         make(chan string, 1),
			Subscriptions: map[string][]protocol.Notifier{},
		}
		_ = board.Subscribe("1", []protocol.Notifier{email})
		_ = board.Subscribe("2", []protocol.Notifier{email})
		assert.Len(t, board.Subscriptions, 2)
	}

	{
		t.Log("Should only add successful subscriptions")
		email, err := notifier.NewNotifier(protocol.Email)
		assert.Nil(t, err)
		board := MessageBoard{
			Id:            "",
			MsgCh:         make(chan string, 1),
			Subscriptions: map[string][]protocol.Notifier{},
		}
		_ = board.Subscribe("1", []protocol.Notifier{email})
		_ = board.Subscribe("2", []protocol.Notifier{})
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
			MsgCh:         make(chan string, 1),
			Subscriptions: map[string][]protocol.Notifier{},
		}
		_ = board.Subscribe("1", []protocol.Notifier{email})
		assert.Len(t, board.Subscriptions, 1)
		board.Unsubscribe("1")
		assert.Len(t, board.Subscriptions, 0)
	}
}
