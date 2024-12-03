package main

import (
	"fmt"
	"log"
	"msg-board/daemon"
	"msg-board/protocol"
	"msg-board/repository"
	"os"
)

func main() {
	err := run()
	if err != nil {
		fmt.Println(err)
	}
}

func run() error {
	repo, err := repository.NewRepository(protocol.MemRepo)
	if err != nil {
		return err
	}
	server, err := daemon.NewServer(daemon.Config{
		Addr:       "0.0.0.0:8080",
		Repository: repo,
		Notifiers:  []protocol.NotifyService{protocol.Email, protocol.SMS, protocol.WhatsApp},
	}, log.New(os.Stdout, "HTTP Server:", log.LstdFlags))
	if err != nil {
		return err
	}

	// TODO Gracefully Shutdown
	server.Start()
	return nil
}
