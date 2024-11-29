package board

import (
	"io"
	"msg-board/protocol"
	"msg-board/service/notifier"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewBoard(t *testing.T) {
	t.Run("Should create new board", func(t *testing.T) {
		res := NewBoard("", "")
		assert.NotNil(t, res)
	})
}

func TestSubscribe(t *testing.T) {
	t.Run("Should subscribe to one service", func(t *testing.T) {
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
		err = board.Subscribe(subscription, "")
		assert.Len(t, board.Subscriptions[userId], 1)
		assert.Nil(t, err)
	})

	t.Run("Should subscribe to multiple services", func(t *testing.T) {
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
		err = board.Subscribe(subscription, "")
		assert.Nil(t, err)
		assert.Len(t, board.Subscriptions[userId], 2)
	})

	t.Run("Should error on subscribe without notifier", func(t *testing.T) {
		userId := ""
		board := MessageBoard{
			Id:            userId,
			Subscriptions: map[string][]protocol.Notifier{},
		}
		subscription := protocol.Subscription{
			Subscriber: userId,
			Services:   nil,
		}
		err := board.Subscribe(subscription, "")
		assert.Error(t, err)
	})
}

func TestMultipleSubscribe(t *testing.T) {
	t.Run("Should add two subscriptions", func(t *testing.T) {
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
		err = board.Subscribe(subscription1, "")
		assert.Nil(t, err)
		subscription2 := protocol.Subscription{
			Subscriber: "2",
			Services:   []protocol.Notifier{email},
		}
		err = board.Subscribe(subscription2, "")
		assert.Nil(t, err)

		assert.Len(t, board.Subscriptions, 2)
	})

	t.Run("Should only add sucessful subscriptions", func(t *testing.T) {
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
		err = board.Subscribe(subscription1, "")
		assert.Nil(t, err)
		subscription2 := protocol.Subscription{
			Subscriber: "2",
			Services:   []protocol.Notifier{},
		}
		err = board.Subscribe(subscription2, "")
		assert.Error(t, err)

		assert.Len(t, board.Subscriptions, 1)
	})
}

func TestUnsubscribe(t *testing.T) {
	t.Run("Should delete on unsubscribe", func(t *testing.T) {
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
		err = board.Subscribe(subscription, "")
		assert.Nil(t, err)

		assert.Len(t, board.Subscriptions, 1)
		board.Unsubscribe("1")
		assert.Len(t, board.Subscriptions, 0)
	})
}

func TestNewMessage(t *testing.T) {
	t.Run("Should message one user with multiple notifiers", func(t *testing.T) {
		email, err := notifier.NewNotifier(protocol.Email)
		assert.Nil(t, err)
		sms, err := notifier.NewNotifier(protocol.SMS)
		assert.Nil(t, err)
		board := NewBoard("", "")
		subscription := protocol.Subscription{
			Subscriber: "1",
			Services:   []protocol.Notifier{email, sms},
		}
		err = board.Subscribe(subscription, "")
		assert.Nil(t, err)

		r, w, _ := os.Pipe()
		os.Stdout = w
		board.NewMessage("")
		w.Close()

		bytes, _ := io.ReadAll(r)
		assert.Equal(t, "Email: User 1 \nSMS: User 1 \n", string(bytes))
	})

	t.Run("Should message two users", func(t *testing.T) {
		email, err := notifier.NewNotifier(protocol.Email)
		assert.Nil(t, err)
		board := NewBoard("", "")
		subscription1 := protocol.Subscription{
			Subscriber: "1",
			Services:   []protocol.Notifier{email},
		}
		subscription2 := protocol.Subscription{
			Subscriber: "2",
			Services:   []protocol.Notifier{email},
		}
		err = board.Subscribe(subscription1, "")
		assert.Nil(t, err)
		err = board.Subscribe(subscription2, "")
		assert.Nil(t, err)

		r, w, _ := os.Pipe()
		os.Stdout = w
		board.NewMessage("")
		w.Close()

		bytes, _ := io.ReadAll(r)
		assert.Equal(t, "Email: User 1 \nEmail: User 2 \n", string(bytes))
	})
}
