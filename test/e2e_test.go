package test

import (
	"msg-board/protocol"
	"msg-board/service/board"
	"msg-board/service/notifier"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestE2E(t *testing.T) {
	t.Run("Should stop getting notifications after unsubscribe", func(t *testing.T) {
		pw := ""
		board := board.NewBoard("public", pw)

		whatsApp, _ := notifier.NewNotifiers([]protocol.NotifyService{protocol.WhatsApp})
		err := board.Subscribe(protocol.Subscription{
			Subscriber: "1",
			Services:   whatsApp,
		}, pw)
		assert.Nil(t, err)

		msg := BoardGetMessage(board)
		assert.Contains(t, msg, "WhatsApp: Board public, User 1 test")

		board.Unsubscribe("1")
		msg = BoardGetMessage(board)
		assert.NotContains(t, msg, "WhatsApp: Board public, User 1 test")
	})
	t.Run("Should subscribe users to board and get notified on message", func(t *testing.T) {
		pw := ""
		board := board.NewBoard("public", pw)

		emailSms, _ := notifier.NewNotifiers([]protocol.NotifyService{protocol.Email, protocol.SMS})
		whatsApp, _ := notifier.NewNotifiers([]protocol.NotifyService{protocol.WhatsApp})

		err := board.Subscribe(protocol.Subscription{
			Subscriber: "1",
			Services:   emailSms,
		}, pw)
		assert.Nil(t, err)

		err = board.Subscribe(protocol.Subscription{
			Subscriber: "2",
			Services:   whatsApp,
		}, pw)
		assert.Nil(t, err)

		msg := BoardGetMessage(board)
		assert.Contains(t, msg, "Email: Board public, User 1 test")
		assert.Contains(t, msg, "SMS: Board public, User 1 test")
		assert.Contains(t, msg, "WhatsApp: Board public, User 2 test")
	})

	t.Run("Should subscribe user to multiple boards and get notified on respective board message", func(t *testing.T) {
		publicPw := ""
		privatePw := "test"
		public := board.NewBoard("public", publicPw)
		private := board.NewBoard("private", privatePw)

		email, _ := notifier.NewNotifiers([]protocol.NotifyService{protocol.Email})
		whatsApp, _ := notifier.NewNotifiers([]protocol.NotifyService{protocol.WhatsApp})

		err := public.Subscribe(protocol.Subscription{
			Subscriber: "1",
			Services:   email,
		}, publicPw)
		assert.Nil(t, err)

		err = private.Subscribe(protocol.Subscription{
			Subscriber: "1",
			Services:   whatsApp,
		}, privatePw)
		assert.Nil(t, err)

		msg := BoardGetMessage(public)
		assert.Contains(t, msg, "Email: Board public, User 1 test")
		assert.NotContains(t, msg, "WhatsApp: Board private, User 1 test")

		msg = BoardGetMessage(private)
		assert.NotContains(t, msg, "Email: Board public, User 1 test")
		assert.Contains(t, msg, "WhatsApp: Board private, User 1 test")
	})
}
