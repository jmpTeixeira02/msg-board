package board

import (
	"msg-board/protocol"
	"msg-board/repository/memory"
	"msg-board/test"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewService(t *testing.T) {
	t.Run("Should create new service", func(t *testing.T) {
		s, err := NewService(memory.NewBd(), protocol.Email)
		assert.NotNil(t, s)
		assert.Nil(t, err)
	})
}

func TestSubscribe(t *testing.T) {
	t.Run("Should error on subscribe without notifier", func(t *testing.T) {
		userId := ""
		board := ""

		s, err := NewService(memory.NewBd(), protocol.Email)
		assert.NotNil(t, s)
		assert.Nil(t, err)

		s.repo.AddPublicBoard(board)

		assert.Nil(t, err)

		subscription := protocol.Subscription{
			Subscriber: protocol.Subscribing{
				User:     userId,
				Services: []protocol.NotifyService{},
			},
			Publisher: board,
		}

		err = s.Subscribe(subscription, "")
		assert.Error(t, err)
	})

	t.Run("Should get error subscribing to private board with wrong password", func(t *testing.T) {
		userId := ""
		board := ""

		s, err := NewService(memory.NewBd(), protocol.Email)
		assert.NotNil(t, s)
		assert.Nil(t, err)

		s.repo.AddPrivateBoard(board, "a")

		assert.Nil(t, err)

		subscription := protocol.Subscription{
			Subscriber: protocol.Subscribing{
				User:     userId,
				Services: []protocol.NotifyService{protocol.Email},
			},
			Publisher: board,
		}

		err = s.Subscribe(subscription, "b")
		assert.NotNil(t, err)
	})

	t.Run("Should subscribe to private board", func(t *testing.T) {
		userId := ""
		board := ""

		s, err := NewService(memory.NewBd(), protocol.Email)
		assert.NotNil(t, s)
		assert.Nil(t, err)

		s.repo.AddPrivateBoard(board, "b")

		assert.Nil(t, err)

		subscription := protocol.Subscription{
			Subscriber: protocol.Subscribing{
				User:     userId,
				Services: []protocol.NotifyService{protocol.Email},
			},
			Publisher: board,
		}

		err = s.Subscribe(subscription, "b")
		assert.Nil(t, err)
	})

	t.Run("Should subscribe to public board", func(t *testing.T) {
		userId := ""
		board := ""

		s, err := NewService(memory.NewBd(), protocol.Email)
		assert.NotNil(t, s)
		assert.Nil(t, err)

		s.repo.AddPublicBoard(board)

		assert.Nil(t, err)

		subscription := protocol.Subscription{
			Subscriber: protocol.Subscribing{
				User:     userId,
				Services: []protocol.NotifyService{protocol.Email},
			},
			Publisher: board,
		}

		err = s.Subscribe(subscription, "b")
		assert.Nil(t, err)
	})
}

// TODO Instead of checking msg content there should be a spyOn the send method
func TestNewMessage(t *testing.T) {
	t.Run("Should message one user with multiple notifiers", func(t *testing.T) {
		board := "test"
		s, err := NewService(memory.NewBd(), protocol.Email, protocol.SMS)
		assert.NotNil(t, s)
		assert.Nil(t, err)

		s.repo.AddPublicBoard(board)
		subscription := protocol.Subscription{
			Subscriber: protocol.Subscribing{
				User:     "1",
				Services: []protocol.NotifyService{protocol.Email, protocol.SMS},
			},
			Publisher: board,
		}

		err = s.Subscribe(subscription, "")
		assert.Nil(t, err)

		msg := test.SendMessageAux(func() {
			s.NewMessage(board, "test")
		})
		assert.Contains(t, msg, "User: 1 Email: test")
		assert.Contains(t, msg, "User: 1 SMS: test")
	})

	t.Run("Should message two users", func(t *testing.T) {
		board := "test"
		s, err := NewService(memory.NewBd(), protocol.Email, protocol.SMS)
		assert.NotNil(t, s)
		assert.Nil(t, err)

		s.repo.AddPublicBoard(board)
		subscription := protocol.Subscription{
			Subscriber: protocol.Subscribing{
				User:     "1",
				Services: []protocol.NotifyService{protocol.Email},
			},
			Publisher: board,
		}
		err = s.Subscribe(subscription, "")
		assert.Nil(t, err)

		subscription = protocol.Subscription{
			Subscriber: protocol.Subscribing{
				User:     "2",
				Services: []protocol.NotifyService{protocol.Email},
			},
			Publisher: board,
		}
		err = s.Subscribe(subscription, "")
		assert.Nil(t, err)

		msg := test.SendMessageAux(func() {
			s.NewMessage(board, "test")
		})
		assert.Contains(t, msg, "User: 1 Email: test")
		assert.Contains(t, msg, "User: 2 Email: test")
	})
}
