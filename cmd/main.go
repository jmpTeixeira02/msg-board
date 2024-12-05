package main

import (
	"context"
	"fmt"
	"log"
	"msg-board/daemon"
	"msg-board/protocol"
	"msg-board/repository"
	"os"
	"os/signal"
	"syscall"
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

	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	go server.Start()

	// Gracefully Shutdown
	<-ctx.Done()
	fmt.Println("Shutting down the server")
	return server.Shutdown(ctx)
}
