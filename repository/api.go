package repository

import (
	"fmt"
	"msg-board/protocol"
	"msg-board/repository/memory"
)

// TODO Create a repository using a BD such as SQLite and ORM data models
func NewRepository(repo protocol.Repositories) (protocol.Repo, error) {
	switch protocol.Repositories(repo) {
	case protocol.MemRepo:
		return memory.NewBd(), nil
	default:
		return nil, fmt.Errorf("unimplemented notifier")
	}
}
