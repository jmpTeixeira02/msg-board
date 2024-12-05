package daemon

import (
	"encoding/json"
	"msg-board/daemon/api/generated"
	"msg-board/protocol"
	"msg-board/util"
	"net/http"
)

// (POST /board)
func (s *Server) CreateBoard(w http.ResponseWriter, r *http.Request) {
	newBoardDto := generated.NewBoard{}
	err := json.NewDecoder(r.Body).Decode(&newBoardDto)
	if err != nil {
		util.WriteJsonResponse(w, http.StatusBadRequest, err)
		return
	}
	board := s.Service.NewBoard(newBoardDto.Board, newBoardDto.Password)
	util.WriteJsonResponse(w, http.StatusOK, BoardToBoardDto(board))
}

// (POST /board/{board})
func (s *Server) SendMessage(w http.ResponseWriter, r *http.Request, board string) {
	messageDto := generated.Message{}
	err := json.NewDecoder(r.Body).Decode(&messageDto)
	if err != nil {
		util.WriteJsonResponse(w, http.StatusBadRequest, err)
		return
	}
	s.Service.NewMessage(board, messageDto.Msg)
	util.WriteJsonResponse(w, http.StatusOK, messageDto)
}

// (DELETE /subscription/{board})
func (s *Server) UnsubscribeBoard(w http.ResponseWriter, r *http.Request, board string, user string) {
	unsub := s.Service.Unsubscribe(board, user)
	util.WriteJsonResponse(w, http.StatusOK, UnsubscribeToUnsubscribeDto(unsub))
}

// (POST /subscription/{board})
func (s *Server) SubscribeBoard(w http.ResponseWriter, r *http.Request, board string) {
	subscribingDto := generated.Subscribe{}
	err := json.NewDecoder(r.Body).Decode(&subscribingDto)
	if err != nil {
		util.WriteJsonResponse(w, http.StatusBadRequest, Error{
			Error: err.Error(),
		})
		return
	}
	err = s.Service.Subscribe(protocol.Subscription{
		Subscriber: SubscribingDtoToSubscribing(subscribingDto),
		Publisher:  board,
	}, subscribingDto.Password)
	if err != nil {
		util.WriteJsonResponse(w, http.StatusBadRequest, Error{
			Error: err.Error(),
		})
		return
	}
	util.WriteJsonResponse(w, http.StatusOK, generated.Subscription{
		Board:     board,
		Notifiers: subscribingDto.Notifiers,
		Password:  subscribingDto.Password,
		User:      subscribingDto.User,
	})
}
