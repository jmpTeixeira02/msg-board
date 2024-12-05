package board

import (
	"msg-board/protocol"
	"msg-board/repository/memory"
	"msg-board/util"
	"testing"

	"github.com/stretchr/testify/assert"
)

func initTest() BoardService {
	s, err := NewService(memory.NewBd(), protocol.Email, protocol.SMS, protocol.WhatsApp)
	if err != nil {
		panic(err)
	}
	return s
}

func TestNewService(t *testing.T) {
	t.Run("Should create new service", func(t *testing.T) {
		s, err := NewService(memory.NewBd(), protocol.Email)
		assert.NotNil(t, s)
		assert.Nil(t, err)
	})
}

func TestNewBoard(t *testing.T) {
	t.Run("Should create private board", func(t *testing.T) {
		s := initTest()

		boardNae := "board"
		pw := "a"
		board := s.NewBoard(boardNae, &pw)
		assert.True(t, board.Private)
		assert.Equal(t, "a", board.Password)
		assert.Len(t, board.Subscriptions, 0)
	})

	t.Run("Should create public board", func(t *testing.T) {
		s := initTest()

		boardName := "board"
		pw := ""
		board := s.NewBoard(boardName, &pw)
		assert.False(t, board.Private)
		assert.Equal(t, "", board.Password)
		assert.Len(t, board.Subscriptions, 0)
	})
}

func TestSubscribe(t *testing.T) {
	t.Run("Should error on subscribe without notifier", func(t *testing.T) {
		userId := ""
		board := ""

		s := initTest()

		pw := ""
		s.NewBoard(board, &pw)
		subscription := protocol.Subscription{
			Subscriber: protocol.Subscribing{
				User:     userId,
				Services: []protocol.NotifyService{},
			},
			Publisher: board,
		}

		err := s.Subscribe(subscription, &pw)
		assert.Error(t, err)
	})

	t.Run("Should get error subscribing to private board with wrong password", func(t *testing.T) {
		userId := ""
		board := ""

		s := initTest()

		pw := "a"
		s.NewBoard(board, &pw)

		subscription := protocol.Subscription{
			Subscriber: protocol.Subscribing{
				User:     userId,
				Services: []protocol.NotifyService{protocol.Email},
			},
			Publisher: board,
		}

		pwB := "b"
		err := s.Subscribe(subscription, &pwB)
		assert.NotNil(t, err)
	})

	t.Run("Should subscribe to private board", func(t *testing.T) {
		userId := ""
		board := ""

		s := initTest()

		pw := "b"
		s.NewBoard(board, &pw)

		subscription := protocol.Subscription{
			Subscriber: protocol.Subscribing{
				User:     userId,
				Services: []protocol.NotifyService{protocol.Email},
			},
			Publisher: board,
		}

		err := s.Subscribe(subscription, &pw)
		assert.Nil(t, err)
	})

	t.Run("Should subscribe to public board", func(t *testing.T) {
		userId := ""
		board := ""

		s := initTest()

		pw := ""
		s.NewBoard(board, &pw)
		subscription := protocol.Subscription{
			Subscriber: protocol.Subscribing{
				User:     userId,
				Services: []protocol.NotifyService{protocol.Email},
			},
			Publisher: board,
		}

		err := s.Subscribe(subscription, nil)
		assert.Nil(t, err)
	})
}

// TODO Instead of checking msg content there should be a spyOn the send method
func TestNewMessage(t *testing.T) {
	t.Run("Should message one user with multiple notifiers", func(t *testing.T) {
		board := "test"
		s := initTest()

		pw := ""
		s.NewBoard(board, &pw)
		subscription := protocol.Subscription{
			Subscriber: protocol.Subscribing{
				User:     "1",
				Services: []protocol.NotifyService{protocol.Email, protocol.SMS},
			},
			Publisher: board,
		}

		err := s.Subscribe(subscription, nil)
		assert.Nil(t, err)

		msg := util.CaptureStdOutput(func() {
			s.NewMessage(board, "test")
		})
		assert.Contains(t, msg, "User: 1 Email: test")
		assert.Contains(t, msg, "User: 1 SMS: test")
	})

	t.Run("Should message two users", func(t *testing.T) {
		board := "test"
		s := initTest()

		pw := ""
		s.NewBoard(board, &pw)
		subscription := protocol.Subscription{
			Subscriber: protocol.Subscribing{
				User:     "1",
				Services: []protocol.NotifyService{protocol.Email},
			},
			Publisher: board,
		}
		err := s.Subscribe(subscription, nil)
		assert.Nil(t, err)

		subscription = protocol.Subscription{
			Subscriber: protocol.Subscribing{
				User:     "2",
				Services: []protocol.NotifyService{protocol.Email},
			},
			Publisher: board,
		}
		err = s.Subscribe(subscription, nil)
		assert.Nil(t, err)

		msg := util.CaptureStdOutput(func() {
			s.NewMessage(board, "test")
		})
		assert.Contains(t, msg, "User: 1 Email: test")
		assert.Contains(t, msg, "User: 2 Email: test")
	})
}
