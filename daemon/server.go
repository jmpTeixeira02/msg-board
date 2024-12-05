package daemon

import (
	"context"
	"log"
	api "msg-board/daemon/api/generated"
	"msg-board/protocol"
	"msg-board/service/board"
	"net/http"
)

type Config struct {
	Addr       string
	Repository protocol.Repo
	Notifiers  []protocol.NotifyService
}

type Server struct {
	Addr     string
	Log      *log.Logger
	Service  board.BoardService
	Repo     protocol.Repo
	Shutdown func(context.Context) error
}

func NewServer(config Config, log *log.Logger) (Server, error) {
	service, err := board.NewService(config.Repository, config.Notifiers...)
	if err != nil {
		return Server{}, err
	}
	return Server{
		Addr:    config.Addr,
		Log:     log,
		Service: service,
		Repo:    config.Repository,
	}, nil
}

func (s *Server) Start() {
	r := http.NewServeMux()

	h := api.HandlerFromMux(s, r)

	server := &http.Server{
		Handler: h,
		Addr:    s.Addr,
	}

	s.Log.Printf("Listening on %s\n", s.Addr)
	s.Shutdown = server.Shutdown

	s.Log.Fatal(server.ListenAndServe())
}
