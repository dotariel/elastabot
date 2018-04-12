package slack

import (
	"errors"
	"testing"
	"time"

	"github.com/dotariel/elastabot/elastabot"
	"github.com/stretchr/testify/assert"
)

var (
	server = newMockServer()
	client = newMockClient()
)

func TestNewClient(t *testing.T) {
	c := New()

	go func() { c.connected <- true }()
	<-c.connected

	go func() { c.shutdown <- true }()
	<-c.shutdown
}

func TestParseCommand(t *testing.T) {
	testCases := []struct {
		message  string
		expected elastabot.Command
		err      error
	}{
		{
			message:  "nomatch",
			expected: nil,
			err:      errors.New("unknown command"),
		},
		{
			message: "!ack",
			expected: &elastabot.Ack{
				Alert:    "",
				Duration: 0,
				Triage:   nil,
			},
			err: nil,
		},
		{
			message: "!ack Acme Flatline Alert",
			expected: &elastabot.Ack{
				Alert:    "Acme Flatline Alert",
				Duration: 0,
				Triage:   nil,
			},
			err: nil,
		},
		{
			message: "!ack Acme?",
			expected: &elastabot.Ack{
				Alert:    "Acme",
				Duration: 0,
				Triage: &elastabot.Triage{
					Topic:   "",
					Handler: nil,
				},
			},
			err: nil,
		},
		{
			message: "!ack Acme|10",
			expected: &elastabot.Ack{
				Alert:    "Acme",
				Duration: 10,
				Triage:   nil,
			},
			err: nil,
		},
	}

	for _, tt := range testCases {
		resp, err := client.parseCommand(tt.message)
		assert.Equal(t, tt.err, err)
		assert.ObjectsAreEqual(tt.expected, resp)
	}

}

// func TestAck(t *testing.T) {
// 	testCases := []struct {
// 		message  string
// 		expected string
// 	}{
// 		{
// 			message:  "!ack",
// 			expected: `{"id":1,"channel":"_channel","text":"Acknowledged alert *foo* until XXX","type":"message"}`,
// 		},
// 		{
// 			message:  "!ack ?",
// 			expected: `{"id":2,"channel":"_channel","text":"Acknowledged alert *foo* until XXX","type":"message"}`,
// 		},
// 		{
// 			message:  "!ack|1m",
// 			expected: `{"id":3,"channel":"_channel","text":"Acknowledged alert *foo* until XXX","type":"message"}`,
// 		},
// 		{
// 			message:  "!ack help",
// 			expected: `{"id":4,"channel":"_channel","text":"Acknowledged alert *foo* until XXX","type":"message"}`,
// 		},
// 	}

// 	go client.Start()
// 	defer client.Stop()
// 	<-client.connected

// 	for _, tt := range testCases {
// 		assert.JSONEq(t, tt.expected, sendAndWaitForEvent(client, server, tt.message, 50))
// 	}
// }

func TestShutdown(t *testing.T) {
	go client.Start()
	<-client.connected
	client.Stop()

	assert.Equal(t, "", sendAndWaitForEvent(client, server, "timeout", 50))
}

func waitForEvent(c *Client, s *MockServer, timeoutMs time.Duration) string {
	select {
	case msg := <-s.responses:
		return msg
	case <-time.After(timeoutMs * time.Millisecond):
		return ""
	}
}

func sendAndWaitForEvent(c *Client, s *MockServer, e string, timeoutMs time.Duration) string {
	c.rtm.IncomingEvents <- newMessageEvent(e)
	select {
	case msg := <-s.responses:
		return msg
	case <-time.After(timeoutMs * time.Millisecond):
		return ""
	}
}
