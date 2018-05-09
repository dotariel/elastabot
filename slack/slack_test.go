package slack

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/dotariel/elastabot/elastabot"
	slk "github.com/nlopes/slack"
	"github.com/stretchr/testify/assert"
)

var (
	server = newMockServer()
	client = newMockClient()
)

type FakeCommandParser struct{}

func (p FakeCommandParser) Parse(s string) (elastabot.Command, error) {
	return FakeCommand{}, nil
}

type FakeCommand struct{}

func (c FakeCommand) Execute() (string, error) {
	return "ok", nil
}

func TestNewClient(t *testing.T) {
	c := New()

	go func() { c.connected <- true }()
	<-c.connected

	go func() { c.shutdown <- true }()
	<-c.shutdown
}

func TestHandleMessage(t *testing.T) {
	go client.Start()
	defer client.Stop()
	<-client.connected

	event := &slk.MessageEvent{Msg: slk.Msg{Text: "foo"}}
	client.handleMessage(FakeCommandParser{}, event)

	var msg slk.Msg
	resp := waitForEvent(client, server, 50)
	json.Unmarshal([]byte(resp), &msg)

	assert.Equal(t, msg.Text, "ok")
}

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
	return waitForEvent(c, s, timeoutMs)
}
