package memory

import (
	"errors"
	"msg-board/protocol"
)

type messageBoard struct {
	Subscriptions map[string][]protocol.NotifyService // userId -> []notifiers
	Private       bool
	Password      string
	Name          string
}

type DB struct {
	boards map[string]messageBoard
}

// TODO
// This BD is not concurrent safe
// It's possible to override Subscriptions when adding
func NewBd() protocol.Repo {
	return &DB{
		boards: map[string]messageBoard{},
	}
}

func (db *DB) AddPublicBoard(id string) protocol.Board {
	res := messageBoard{
		Subscriptions: map[string][]protocol.NotifyService{},
		Private:       false,
		Password:      "",
		Name:          id,
	}

	db.boards[id] = res
	return BoardDAOToBoard(res)
}

func (db *DB) AddPrivateBoard(id string, pw string) protocol.Board {
	res := messageBoard{
		Subscriptions: map[string][]protocol.NotifyService{},
		Private:       true,
		Password:      pw,
		Name:          id,
	}
	db.boards[id] = res
	return BoardDAOToBoard(res)
}

func (db *DB) IsPrivateBoard(board string) (bool, string, error) {
	b, ok := db.boards[board]
	if !ok {
		return false, "", errors.New("board does not exist")
	}
	return b.Private, b.Password, nil
}

func (db *DB) Subscribe(board string, subscriber protocol.Subscribing) error {
	_, ok := db.boards[board]
	if !ok {
		return errors.New("board does not exist")
	}
	db.boards[board].Subscriptions[subscriber.User] = subscriber.Services
	return nil
}

func (db *DB) Unsubscribe(board string, user string) protocol.Unsubscribe {
	services := db.boards[board].Subscriptions[user]
	delete(db.boards[board].Subscriptions, user)
	return protocol.Unsubscribe{
		Board:     board,
		User:      user,
		Notifiers: services,
	}
}

func (db *DB) GetSubscribers(board string) []protocol.Subscribing {
	subs := db.boards[board].Subscriptions
	res := make([]protocol.Subscribing, len(subs))
	idx := 0
	for user, notifiers := range subs {
		res[idx] = protocol.Subscribing{
			User:     user,
			Services: notifiers,
		}
		idx++
	}
	return res
}
