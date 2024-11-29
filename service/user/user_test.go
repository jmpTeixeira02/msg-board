package user

import (
	"errors"
	"msg-board/protocol"
	"msg-board/service/notifier"
	"reflect"
	"testing"
)

func TestNewUser(t *testing.T) {
	tests := []struct {
		testName string
		expected User
	}{
		{
			testName: "Should create user",
			expected: User{
				Boards: map[string][]protocol.Notifier{},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			user := NewUser()
			if !reflect.DeepEqual(user, tt.expected) {
				t.Errorf("Got %+v, expected %+v", user, tt.expected)
			}
		})
	}
}

func TestSubscribe(t *testing.T) {
	emailN, _ := notifier.NewNotifier(protocol.Email)
	smsN, _ := notifier.NewNotifier(protocol.SMS)
	tests := []struct {
		testName  string
		notifiers []protocol.Notifier
		expected  error
	}{
		{
			testName:  "Should subscribe with one notifier ",
			notifiers: []protocol.Notifier{emailN},
			expected:  nil,
		},
		{
			testName:  "Should subscribe with multiple notifiers",
			notifiers: []protocol.Notifier{emailN, smsN},
			expected:  nil,
		},
		{
			testName:  "Should error on subscription with no notifiers",
			notifiers: []protocol.Notifier{},
			expected:  errors.New("a subscription must have notifiers"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			user := NewUser()
			err := user.Subscribe("", tt.notifiers)
			if err != nil && err.Error() != tt.expected.Error() {
				t.Errorf("Got %+v, expected %+v", err, tt.expected)
			}

			if tt.expected != nil && len(user.Boards) != 0 {
				t.Errorf("Got %+v, expected %+v", len(user.Boards), 0)
			} else if tt.expected == nil && len(user.Boards[""]) != len(tt.notifiers) {
				t.Errorf("Got %+v, expected %+v", len(user.Boards), len(tt.notifiers))
			}
		})
	}
}
