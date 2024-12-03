package test

import (
	"msg-board/protocol"
	"msg-board/repository"
	"msg-board/service/board"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestE2E(t *testing.T) {
	t.Run("Should stop getting notifications after unsubscribe", func(t *testing.T) {
		repo, err := repository.NewRepository(protocol.MemRepo)
		assert.Nil(t, err)
		s, err := board.NewService(repo, protocol.WhatsApp)
		assert.Nil(t, err)

		board := "public"
		repo.AddPublicBoard(board)

		err = s.Subscribe(protocol.Subscription{
			Subscriber: protocol.Subscribing{
				User:     "1",
				Services: []protocol.NotifyService{protocol.WhatsApp},
			},
			Publisher: board,
		}, "")
		assert.Nil(t, err)

		msg := SendMessageAux(func() {
			s.NewMessage(board, "test")
		})
		assert.Contains(t, msg, "User: 1 WhatsApp: public")

		s.Unsubscribe(board, "1")
		msg = SendMessageAux(func() {
			s.NewMessage(board, "test")
		})
		assert.NotContains(t, msg, "WhatsApp: Board public, User 1 test")
	})

	t.Run("Should subscribe users to board and get notified on message", func(t *testing.T) {
		repo, err := repository.NewRepository(protocol.MemRepo)
		assert.Nil(t, err)
		s, err := board.NewService(repo, protocol.WhatsApp, protocol.Email, protocol.SMS)
		assert.Nil(t, err)

		board := "public"
		repo.AddPublicBoard(board)

		err = s.Subscribe(protocol.Subscription{
			Subscriber: protocol.Subscribing{
				User:     "1",
				Services: []protocol.NotifyService{protocol.Email, protocol.SMS},
			},
			Publisher: board,
		}, "")
		assert.Nil(t, err)

		err = s.Subscribe(protocol.Subscription{
			Subscriber: protocol.Subscribing{
				User:     "2",
				Services: []protocol.NotifyService{protocol.WhatsApp},
			},
			Publisher: board,
		}, "")
		assert.Nil(t, err)

		msg := SendMessageAux(func() {
			s.NewMessage(board, "test")
		})
		assert.Contains(t, msg, "User: 1 Email: public")
		assert.Contains(t, msg, "User: 1 SMS: public")
		assert.Contains(t, msg, "User: 2 WhatsApp: public")
	})

	t.Run("Should subscribe user to multiple boards and get notified on respective board message", func(t *testing.T) {
		repo, err := repository.NewRepository(protocol.MemRepo)
		assert.Nil(t, err)
		s, err := board.NewService(repo, protocol.WhatsApp, protocol.Email)
		assert.Nil(t, err)

		public := "public"
		private := "private"
		pw := "private"
		repo.AddPublicBoard(public)
		repo.AddPrivateBoard(private, pw)

		err = s.Subscribe(protocol.Subscription{
			Subscriber: protocol.Subscribing{
				User:     "1",
				Services: []protocol.NotifyService{protocol.Email},
			},
			Publisher: public,
		}, "")
		assert.Nil(t, err)

		err = s.Subscribe(protocol.Subscription{
			Subscriber: protocol.Subscribing{
				User:     "1",
				Services: []protocol.NotifyService{protocol.WhatsApp},
			},
			Publisher: private,
		}, pw)
		assert.Nil(t, err)

		msg := SendMessageAux(func() {
			s.NewMessage(public, "test")
		})
		assert.Contains(t, msg, "User: 1 Email: public")
		assert.NotContains(t, msg, "User: 1 WhatsApp: private")

		msg = SendMessageAux(func() {
			s.NewMessage(private, "test")
		})
		assert.NotContains(t, msg, "User: 1 Email: public")
		assert.Contains(t, msg, "User: 1 WhatsApp: private")
	})
}
