package board

import (
	"errors"
	"msg-board/protocol"
	"msg-board/service/notifier"
)

type BoardService struct {
	repo      protocol.Repo
	notifiers map[protocol.NotifyService]protocol.Notifier
}

func NewService(repo protocol.Repo, notifierServices ...protocol.NotifyService) (BoardService, error) {
	notifierMap := make(map[protocol.NotifyService]protocol.Notifier)
	for i := range notifierServices {
		n, err := notifier.NewNotifier(notifierServices[i])
		if err != nil {
			return BoardService{}, err
		}
		notifierMap[notifierServices[i]] = n
	}
	return BoardService{
		repo:      repo,
		notifiers: notifierMap,
	}, nil
}

func (s *BoardService) Subscribe(subscription protocol.Subscription, pw string) error {
	isPrivate, boardPw, err := s.repo.IsPrivateBoard(subscription.Publisher)
	if err != nil {
		return err
	}
	if isPrivate && pw != boardPw {
		return errors.New("invalid password")
	}
	if len(subscription.Subscriber.Services) < 1 {
		return errors.New("subscription must have notifiers")
	}
	return s.repo.Subscribe(subscription.Publisher, subscription.Subscriber)
}

func (s *BoardService) Unsubscribe(board string, user string) {
	s.repo.Unsubscribe(board, user)
}

func (s *BoardService) NewMessage(board string, msg string) {
	subs := s.repo.GetSubscribers(board)
	for i := range subs {
		for j := range subs[i].Services {
			s.notifiers[subs[i].Services[j]].Send(subs[i].User, board)
		}
	}
}
