package test

import (
	"io"
	"msg-board/protocol"
	"os"
)

func SendMessageAux(sendMessage func()) string {
	r, w, _ := os.Pipe()
	os.Stdout = w

	sendMessage()
	w.Close()
	bytes, _ := io.ReadAll(r)

	return string(bytes)
}

func NotifierGetMessage(user string, notifier protocol.Notifier) string {
	r, w, _ := os.Pipe()
	os.Stdout = w
	notifier.Send(user, "test")
	w.Close()

	bytes, _ := io.ReadAll(r)
	return string(bytes)
}
