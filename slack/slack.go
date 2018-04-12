package slack

import (
	"errors"
	"regexp"
	"strings"

	"github.com/dotariel/elastabot/elastabot"
	"github.com/nlopes/slack"
	log "github.com/sirupsen/logrus"
)

const (
	cmdPrefix = "!"
	token     = "TOKEN"
)

// Client wraps a Slack client with channels for synchronizing state
type Client struct {
	rtm       *slack.RTM
	channel   string
	connected chan bool
	shutdown  chan bool
}

// New creates a Slack client
func New() *Client {
	return &Client{
		rtm:       slack.New(token).NewRTM(),
		connected: make(chan bool, 1),
		shutdown:  make(chan bool, 1),
	}
}

// Start opens a connection to Slack and begins listening to events
func (c *Client) Start() {
	go c.rtm.ManageConnection()

	for {
		select {
		case <-c.shutdown:
			return
		case msg := <-c.rtm.IncomingEvents:
			switch ev := msg.Data.(type) {

			case *slack.ConnectedEvent:
				log.Info("connected to Slack")
				c.connected <- true

			case *slack.MessageEvent:
				c.handleEvent(ev)

			default:
			}
		}
	}
}

func (c *Client) handleEvent(event *slack.MessageEvent) {
	command, err := c.parseCommand(event.Msg.Text)
	if err != nil {
		c.respond(err.Error())
		return
	}

	response, err := command.Execute()
	if err != nil {
		c.respond(err.Error())
		return
	}

	c.respond(response)
}

func (c *Client) parseCommand(cmd string) (elastabot.Command, error) {
	re, _ := regexp.Compile(`^!(?P<command>ack|triage|help) ?(?P<alert>[^|?]+)? ?[|]?(?P<duration>[\d]+[smhdw]?)?`)

	keys := re.SubexpNames()
	vals := re.FindAllStringSubmatch(cmd, -1)

	if len(vals) == 0 {
		return nil, errors.New("unknown command")
	}

	md := map[string]string{}
	for i, n := range vals[0] {
		md[keys[i]] = n
	}

	hasTriage := strings.Contains(cmd, "?")

	switch md["command"] {
	case "ack":
		var triage *elastabot.Triage
		var alert = md["alert"]
		var duration = md["duration"]

		if hasTriage {
			triage = &elastabot.Triage{
				Topic:   alert,
				Handler: nil,
			}
		}

		return elastabot.NewAck(alert, duration, triage), nil
	}

	return nil, nil
}

func (c *Client) respond(message string) {
	c.rtm.SendMessage(c.rtm.NewOutgoingMessage(message, c.channel))
}

// Stop shuts down the client
func (c *Client) Stop() {
	log.Info("shutdown signal received")
	c.shutdown <- true
}
