package memory

import (
	"fmt"
	"msg-board/protocol"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	board = "board"
	pw    = "private"
)

func TestNewDb(t *testing.T) {
	t.Run("Should create new db", func(t *testing.T) {
		repo := NewBd()
		assert.NotNil(t, repo)
	})
}

func TestAddBoard(t *testing.T) {
	t.Run("Should create new public board", func(t *testing.T) {
		repo := &DB{
			boards: map[string]messageBoard{},
		}
		assert.NotNil(t, repo)
		repo.AddPublicBoard(board)

		assert.Len(t, repo.boards, 1)
		assert.False(t, repo.boards[board].Private)
	})

	t.Run("Should create new private board", func(t *testing.T) {
		repo := &DB{
			boards: map[string]messageBoard{},
		}
		assert.NotNil(t, repo)
		repo.AddPrivateBoard(board, pw)

		assert.Len(t, repo.boards, 1)
		assert.True(t, repo.boards[board].Private)
		assert.Equal(t, repo.boards[board].Password, pw)
	})

	t.Run("Should create multiple boards", func(t *testing.T) {
		repo := &DB{
			boards: map[string]messageBoard{},
		}
		assert.NotNil(t, repo)

		repo.AddPrivateBoard(board, pw)
		repo.AddPublicBoard("public")

		assert.Len(t, repo.boards, 2)
	})
}

func TestIsPrivateBoard(t *testing.T) {
	t.Run("Should be true on private board", func(t *testing.T) {
		repo := &DB{
			boards: map[string]messageBoard{},
		}
		assert.NotNil(t, repo)

		repo.AddPrivateBoard(board, pw)
		isPrivate, pw, err := repo.IsPrivateBoard(board)
		assert.True(t, isPrivate)
		assert.Equal(t, repo.boards[board].Password, pw)
		assert.Nil(t, err)
	})

	t.Run("Should be false on public board", func(t *testing.T) {
		repo := &DB{
			boards: map[string]messageBoard{},
		}
		assert.NotNil(t, repo)

		repo.AddPublicBoard(board)
		isPrivate, _, err := repo.IsPrivateBoard(board)
		assert.False(t, isPrivate)
		assert.Nil(t, err)
	})
	t.Run("Should error on non-existing board", func(t *testing.T) {
		repo := &DB{
			boards: map[string]messageBoard{},
		}
		assert.NotNil(t, repo)

		_, _, err := repo.IsPrivateBoard(board)
		assert.NotNil(t, err)
	})
}

func TestSubscribe(t *testing.T) {
	t.Run("Should error on non-existing board", func(t *testing.T) {
		repo := &DB{
			boards: map[string]messageBoard{},
		}
		assert.NotNil(t, repo)

		err := repo.Subscribe(board, protocol.Subscribing{
			User:     "",
			Services: []protocol.NotifyService{},
		})
		assert.NotNil(t, err)
	})

	t.Run("Should subscribe", func(t *testing.T) {
		repo := &DB{
			boards: map[string]messageBoard{},
		}
		assert.NotNil(t, repo)

		user := "user"
		sub := protocol.Subscribing{
			User:     user,
			Services: []protocol.NotifyService{protocol.Email},
		}
		repo.AddPublicBoard(board)
		err := repo.Subscribe(board, sub)
		assert.Nil(t, err)
		assert.Equal(t, sub.Services, repo.boards[board].Subscriptions[sub.User])
	})
}

func TestGetSubscribers(t *testing.T) {
	t.Run("Should get all subscribers", func(t *testing.T) {
		repo := &DB{
			boards: map[string]messageBoard{},
		}
		assert.NotNil(t, repo)

		repo.AddPublicBoard(board)
		sub := protocol.Subscribing{
			User:     "user0",
			Services: []protocol.NotifyService{protocol.Email},
		}
		err := repo.Subscribe(board, sub)
		assert.Nil(t, err)

		sub = protocol.Subscribing{
			User:     "user1",
			Services: []protocol.NotifyService{protocol.WhatsApp, protocol.Email},
		}

		err = repo.Subscribe(board, sub)
		assert.Nil(t, err)

		subs := repo.GetSubscribers(board)

		// GetSubscribers return from a map so order is not always the same
		sort.Slice(subs, func(i, j int) bool {
			return subs[i].User < subs[j].User
		})

		for i := range subs {
			u := fmt.Sprintf("user%d", i)
			assert.Equal(t, u, subs[i].User)
			assert.Equal(t, repo.boards[board].Subscriptions[u], subs[i].Services)
		}
	})
}
