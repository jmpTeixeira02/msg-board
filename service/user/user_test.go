package user

import (
	"msg-board/protocol"
	"msg-board/service/board"
	"msg-board/service/notifier"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	{
		t.Logf("Should create user")
		user := NewUser()
		assert.NotNil(t, user)
	}
}

func TestSubscribe(t *testing.T) {
	emailN, _ := notifier.NewNotifier(protocol.Email)
	smsN, _ := notifier.NewNotifier(protocol.SMS)

	{
		t.Logf("Should subscribe with one notifier")
		notifiers := []protocol.Notifier{emailN}
		user := NewUser()
		err := user.Subscribe(board.NewBoard(""), notifiers)
		assert.Nil(t, err)
		assert.Len(t, user.Boards[""].NotifyServices, len(notifiers))
		assert.Len(t, user.Boards, 1)
	}

	{
		t.Logf("Should subscribe with multiple notifiers")
		notifiers := []protocol.Notifier{emailN, smsN}
		user := NewUser()
		err := user.Subscribe(board.NewBoard(""), notifiers)
		assert.Nil(t, err)
		assert.Len(t, user.Boards[""].NotifyServices, len(notifiers))
		assert.Len(t, user.Boards, 1)
	}

	{
		t.Logf("Should error on subscription with no notifiers")
		notifiers := []protocol.Notifier{}
		user := NewUser()
		err := user.Subscribe(board.NewBoard(""), notifiers)
		assert.Error(t, err)
		assert.Len(t, user.Boards, 0)
	}
}
