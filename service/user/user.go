package user

import (
	"errors"
	"msg-board/protocol"
	"msg-board/service/board"

	"github.com/google/uuid"
)

type User struct {
	Id     string
	Boards map[string]protocol.Subscription // HashSet
}

func NewUser() User {
	return User{
		Id:     uuid.NewString(),
		Boards: map[string]protocol.Subscription{},
	}
}

func (u *User) Subscribe(board board.MessageBoard, services []protocol.Notifier) error {
	if len(services) < 1 {
		return errors.New("a subscription must have notifiers")
	}
	u.Boards[board.Id] = protocol.Subscription{
		NotifyServices: services,
		MsgCh:          board.MsgCh,
	}
	return nil
}

func (u *User) Unsubscribe(board string) {
	delete(u.Boards, board)
}
