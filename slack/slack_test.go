package slack

import (
	"testing"
	"time"

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
