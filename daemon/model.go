package daemon

import (
	"msg-board/daemon/api/generated"
	"msg-board/protocol"
)

type Error struct {
	Error string `json:"error"`
}

func SubscribingDtoToSubscribing(subscribing generated.Subscribe) protocol.Subscribing {
	notifierServices := make([]protocol.NotifyService, len(subscribing.Notifiers))
	for i := range subscribing.Notifiers {
		notifierServices[i] = protocol.NotifyService(subscribing.Notifiers[i])
	}
	return protocol.Subscribing{
		User:     subscribing.User,
		Services: notifierServices,
	}
}

func BoardToBoardDto(board protocol.Board) generated.Board {
	return generated.Board{
		Board:    board.Name,
		Password: board.Password,
		Private:  &board.Private,
	}
}

func UnsubscribeToUnsubscribeDto(unsub protocol.Unsubscribe) generated.Subscription {
	notifiers := make([]string, len(unsub.Notifiers))
	for i := range unsub.Notifiers {
		notifiers[i] = string(unsub.Notifiers[i])
	}
	return generated.Subscription{
		Board:     unsub.Board,
		Notifiers: notifiers,
		Password:  nil,
		User:      unsub.User,
	}
}
