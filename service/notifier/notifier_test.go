package notifier

import (
	"io"
	"msg-board/protocol"
	"os"
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
	t.Run("Should send Email Notification", func(t *testing.T) {
		notifier, _ := NewNotifier(protocol.Email)

		r, w, _ := os.Pipe()
		os.Stdout = w
		notifier.Send("")
		w.Close()

		bytes, _ := io.ReadAll(r)
		assert.Equal(t, "Email: \n", string(bytes))
	})

	t.Run("Should send WhatsApp Notification", func(t *testing.T) {
		notifier, _ := NewNotifier(protocol.WhatsApp)

		r, w, _ := os.Pipe()
		os.Stdout = w
		notifier.Send("")
		w.Close()

		bytes, _ := io.ReadAll(r)
		assert.Equal(t, "WhatsApp: \n", string(bytes))
	})

	t.Run("Should send SMS Notification", func(t *testing.T) {
		notifier, _ := NewNotifier(protocol.SMS)

		r, w, _ := os.Pipe()
		os.Stdout = w
		notifier.Send("")
		w.Close()

		bytes, _ := io.ReadAll(r)
		assert.Equal(t, "SMS: \n", string(bytes))
	})
}
