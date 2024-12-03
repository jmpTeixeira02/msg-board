package daemon

import (
	"msg-board/util"
	"net/http"
)

// TODO Implement
// (POST /board)
func (s *Server) PostBoard(w http.ResponseWriter, r *http.Request) {
	util.WriteReplyJson(w, "TODO")
}

// TODO Implement
// (POST /board/{board})
func (s *Server) PostBoardBoard(w http.ResponseWriter, r *http.Request, board string) {
	util.WriteReplyJson(w, "TODO")
}

// TODO Implement
// (DELETE /subscription/{board})
func (s *Server) DeleteSubscriptionBoard(w http.ResponseWriter, r *http.Request, board string) {
	util.WriteReplyJson(w, "TODO")
}

// TODO Implement
// (POST /subscription/{board})
func (s *Server) PostSubscriptionBoard(w http.ResponseWriter, r *http.Request, board string) {
	util.WriteReplyJson(w, "TODO")
}
