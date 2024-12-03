package notifier

import (
	"msg-board/protocol"
	"msg-board/test"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewNotifier(t *testing.T) {
	t.Run("Should create Email Notifier", func(t *testing.T) {
		notifier, err := NewNotifier(protocol.Email)
		assert.NotNil(t, notifier)
		assert.Nil(t, err)
	})
	t.Run("Should create WhatsApp Notifier", func(t *testing.T) {
		notifier, err := NewNotifier(protocol.WhatsApp)
		assert.NotNil(t, notifier)
		assert.Nil(t, err)
	})
	t.Run("Should create SMS Notifier", func(t *testing.T) {
		notifier, err := NewNotifier(protocol.SMS)
		assert.NotNil(t, notifier)
		assert.Nil(t, err)
	})
	t.Run("Should error on invalid Notifier", func(t *testing.T) {
		notifier, err := NewNotifier(protocol.NotifyService(""))
		assert.Nil(t, notifier)
		assert.NotNil(t, err)
	})
}

func TestSend(t *testing.T) {
	user := "User"
	t.Run("Should send Email Notification", func(t *testing.T) {
		notifier, _ := NewNotifier(protocol.Email)

		msg := test.NotifierGetMessage(user, notifier)
		assert.Contains(t, msg, "User: User Email: test")
	})

	t.Run("Should send WhatsApp Notification", func(t *testing.T) {
		notifier, _ := NewNotifier(protocol.WhatsApp)

		msg := test.NotifierGetMessage(user, notifier)
		assert.Contains(t, msg, "User: User WhatsApp: test")
	})

	t.Run("Should send SMS Notification", func(t *testing.T) {
		notifier, _ := NewNotifier(protocol.SMS)

		msg := test.NotifierGetMessage(user, notifier)
		assert.Contains(t, msg, "User: User SMS: test")
	})
}
