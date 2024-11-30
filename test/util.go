package test

import (
	"io"
	"msg-board/protocol"
	"os"
)

func BoardGetMessage(board protocol.Publisher) string {
	r, w, _ := os.Pipe()
	os.Stdout = w
	board.NewMessage("test")
	w.Close()

	bytes, _ := io.ReadAll(r)
	return string(bytes)
}

func NotifierGetMessage(notifier protocol.Notifier) string {
	r, w, _ := os.Pipe()
	os.Stdout = w
	notifier.Send("test")
	w.Close()

	bytes, _ := io.ReadAll(r)
	return string(bytes)
}
