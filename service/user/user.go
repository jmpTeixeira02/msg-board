package user

import (
	"errors"
	"msg-board/protocol"

	"github.com/google/uuid"
)

type User struct {
	Id     string
	Boards map[string][]protocol.Notifier // HashSet
}

func NewUser() User {
	return User{
		Id:     uuid.NewString(),
		Boards: map[string][]protocol.Notifier{},
	}
}

func (u *User) Subscribe(board string, services []protocol.Notifier) error {
	if len(services) < 1 {
		return errors.New("a subscription must have notifiers")
	}
	u.Boards[board] = services
	return nil
}

func (u *User) Unsubscribe(board string) {
	delete(u.Boards, board)
}
