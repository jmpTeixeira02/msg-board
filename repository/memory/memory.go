package memory

import (
	"errors"
	"msg-board/protocol"
)

type messageBoard struct {
	Subscriptions map[string][]protocol.NotifyService // userId -> []notifiers
	Private       bool
	Password      string
}

type DB struct {
	boards map[string]messageBoard
}

func NewBd() protocol.Repo {
	return &DB{
		boards: map[string]messageBoard{},
	}
}

func (db *DB) AddPublicBoard(id string) {
	db.boards[id] = messageBoard{
		Subscriptions: map[string][]protocol.NotifyService{},
		Private:       false,
		Password:      "",
	}
}

func (db *DB) AddPrivateBoard(id string, pw string) {
	db.boards[id] = messageBoard{
		Subscriptions: map[string][]protocol.NotifyService{},
		Private:       true,
		Password:      pw,
	}
}

func (db *DB) IsPrivateBoard(board string) (bool, string, error) {
	_, ok := db.boards[board]
	if !ok {
		return false, "", errors.New("board does not exist")
	}
	return db.boards[board].Private, db.boards[board].Password, nil
}

func (db *DB) Subscribe(board string, subscriber protocol.Subscribing) error {
	_, ok := db.boards[board]
	if !ok {
		return errors.New("board does not exist")
	}
	db.boards[board].Subscriptions[subscriber.User] = subscriber.Services
	return nil
}

func (db *DB) Unsubscribe(board string, user string) {
	delete(db.boards[board].Subscriptions, user)
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
