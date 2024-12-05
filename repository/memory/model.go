package memory

import (
	"msg-board/protocol"
)

func BoardDAOToBoard(board messageBoard) protocol.Board {
	return protocol.Board{
		Name:          board.Name,
		Subscriptions: board.Subscriptions,
		Private:       board.Private,
		Password:      board.Password,
	}
}
