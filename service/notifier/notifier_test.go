package notifier

import (
	"msg-board/protocol"
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
