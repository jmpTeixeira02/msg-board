package notifier

import (
	"msg-board/protocol"
	"testing"
)

func TestNewNotifier(t *testing.T) {
	tests := []struct {
		testName      string
		notifyService protocol.NotifyService
		expected      protocol.Notifier
	}{
		{
			testName:      "Should create Email Notifier",
			notifyService: protocol.Email,
			expected:      &EmailNotifier{},
		},
		{
			testName:      "Should create WhatsApp Notifier",
			notifyService: protocol.WhatsApp,
			expected:      &WhatsAppNotifier{},
		},
		{
			testName:      "Should create SMS Notifier",
			notifyService: protocol.SMS,
			expected:      &SMSNotifier{},
		},
		{
			testName:      "Should error on unimplemented Notifier",
			notifyService: protocol.NotifyService(""),
			expected:      nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			notifier, _ := NewNotifier(tt.notifyService)
			if notifier != tt.expected {
				t.Errorf("Got %+v, expected %+v", notifier, tt.expected)
			}
		})
	}
}

func TestSend(t *testing.T) {
	tests := []struct {
		testName      string
		notifyService protocol.NotifyService
		expected      error
	}{
		{
			testName:      "Should send email",
			notifyService: protocol.Email,
			expected:      nil,
		},
		{
			testName:      "Should send WhatsApp",
			notifyService: protocol.WhatsApp,
			expected:      nil,
		},
		{
			testName:      "Should send SMS",
			notifyService: protocol.SMS,
			expected:      nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			notifier, _ := NewNotifier(tt.notifyService)
			if notifier.Send("") != tt.expected {
				t.Errorf("Got %+v, expected %+v", notifier, tt.expected)
			}
		})
	}
}
