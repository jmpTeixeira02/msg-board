package repository

import "msg-board/protocol"

type Board struct {
	Subscriptions map[string][]protocol.Notifier // userId -> []notifiers
	Private       bool
	Password      string
}
