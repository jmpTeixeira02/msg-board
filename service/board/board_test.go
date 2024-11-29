package board

import (
	"errors"
	"msg-board/protocol"
	"msg-board/service/notifier"
	"reflect"
	"testing"
)

type errorNotifier struct{}

func (n *errorNotifier) Send(msg string) error {
	return errors.New("err")
}

func TestNewBoard(t *testing.T) {
	tests := []struct {
		testName string
		id       string
		expected MessageBoard
	}{
		{
			testName: "Should create Board",
			id:       "",
			expected: MessageBoard{
				Id:            "",
				Msgs:          []string{},
				Subscriptions: map[string][]protocol.Notifier{},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			res := NewBoard(tt.id)
			if !reflect.DeepEqual(res, tt.expected) {
				t.Errorf("Got %+v, expected %+v", res, tt.expected)
			}
		})
	}
}

func TestSubscribe(t *testing.T) {
	tests := []struct {
		testName string
		sub      protocol.Subscription
		expected error
	}{
		{
			testName: "Should subscribe one service",
			sub: protocol.Subscription{
				UserId:         "",
				NotifyServices: []protocol.NotifyService{protocol.SMS},
			},
			expected: nil,
		},
		{
			testName: "Should subscribe multiple services",
			sub: protocol.Subscription{
				UserId:         "",
				NotifyServices: []protocol.NotifyService{protocol.SMS, protocol.Email},
			},
			expected: nil,
		},
		{
			testName: "Should error on subscribe with unimplemented notifier",
			sub: protocol.Subscription{
				UserId:         "",
				NotifyServices: []protocol.NotifyService{protocol.NotifyService("")},
			},
			expected: errors.New("unimplemented notifier"),
		},
		{
			testName: "Should error on subscribe without notifier",
			sub: protocol.Subscription{
				UserId:         "",
				NotifyServices: []protocol.NotifyService{},
			},
			expected: errors.New("a subscription must have notifiers"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			board := NewBoard("")
			res := board.Subscribe(tt.sub)
			if !reflect.DeepEqual(res, tt.expected) {
				t.Errorf("Got %+v, expected %+v", res, tt.expected)
			}
			if tt.expected == nil && len(board.Subscriptions[tt.sub.UserId]) != len(tt.sub.NotifyServices) {
				t.Errorf("Got %+v, expected %+v", len(board.Subscriptions[tt.sub.UserId]), len(tt.sub.NotifyServices))
			}
		})
	}
}

func TestMultipleSubscribe(t *testing.T) {
	tests := []struct {
		testName string
		subs     []protocol.Subscription
		expected int
	}{
		{
			testName: "Should add two subscriptions",
			subs: []protocol.Subscription{
				{
					UserId:         "1",
					NotifyServices: []protocol.NotifyService{protocol.Email},
				}, {
					UserId:         "2",
					NotifyServices: []protocol.NotifyService{protocol.SMS},
				},
			},
			expected: 2,
		},
		{
			testName: "Should only add successful subscriptions",
			subs: []protocol.Subscription{
				{
					UserId:         "1",
					NotifyServices: []protocol.NotifyService{protocol.Email},
				}, {
					UserId:         "2",
					NotifyServices: []protocol.NotifyService{},
				},
			},
			expected: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			board := NewBoard("")
			for i := range tt.subs {
				_ = board.Subscribe(tt.subs[i])
			}
			if len(board.Subscriptions) != tt.expected {
				t.Errorf("Got %+v, expected %+v", len(board.Subscriptions), tt.expected)
			}
		})
	}
}

func TestUnsubscribe(t *testing.T) {
	tests := []struct {
		testName string
		user     string
	}{
		{
			testName: "Should delete on unsubscribe",
			user:     "1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			board := NewBoard("")
			_ = board.Subscribe(protocol.Subscription{
				UserId:         "1",
				NotifyServices: []protocol.NotifyService{protocol.Email},
			})
			if len(board.Subscriptions) != 1 {
				t.Errorf("Got %+v, expected %+v", len(board.Subscriptions), 1)
			}
			board.Unsubscribe("1")
			if len(board.Subscriptions) != 0 {
				t.Errorf("Got %+v, expected %+v", len(board.Subscriptions), 0)
			}
		})
	}
}

func TestSendMessage(t *testing.T) {
	tests := []struct {
		testName  string
		notifiers []protocol.Notifier
		expected  error
	}{
		{
			testName:  "Should message on notifiers",
			notifiers: []protocol.Notifier{&notifier.EmailNotifier{}},
			expected:  nil,
		},
		{
			testName:  "Should error on unsuccessful notifiers",
			notifiers: []protocol.Notifier{&errorNotifier{}},
			expected:  errors.New("err"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			board := NewBoard("")
			board.Subscriptions["1"] = tt.notifiers
			err := board.SendMessage("")
			if err != nil && err.Error() != tt.expected.Error() {
				t.Errorf("Got %+v, expected %+v", err, tt.expected)
			}
		})
	}
}
