package notifier

import (
	"msg-board/protocol"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewNotifier(t *testing.T) {
	{
		t.Logf("Should create Email Notifier")
		notifier, _ := NewNotifier(protocol.Email)
		assert.NotNil(t, notifier)
	}
	{
		t.Logf("Should create WhatsApp Notifier")
		notifier, _ := NewNotifier(protocol.WhatsApp)
		assert.NotNil(t, notifier)
	}
	{
		t.Logf("Should create SMS Notifier")
		notifier, _ := NewNotifier(protocol.SMS)
		assert.NotNil(t, notifier)
	}
	{
		t.Logf("Should error on invalid Notifier")
		_, err := NewNotifier(protocol.NotifyService(""))
		assert.Error(t, err)
	}
}
