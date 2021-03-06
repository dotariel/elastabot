package slack

import (
	"os"

	"github.com/dotariel/elastabot/elastabot"
	"github.com/nlopes/slack"
	log "github.com/sirupsen/logrus"
)

const cmdPrefix = "!"

var token string

func init() {
	token = os.Getenv("ELASTABOT_SLACK_TOKEN")
}

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
	log.Info("starting bot")

	go c.rtm.ManageConnection()

	for {
		select {
		case <-c.shutdown:
			return
		case msg := <-c.rtm.IncomingEvents:
			switch ev := msg.Data.(type) {
			case *slack.ConnectingEvent:
				log.Info("connecting to Slack")
			case *slack.ConnectionErrorEvent:
				log.Error("could not connect to Slack")
			case *slack.ConnectedEvent:
				log.Info("connected to Slack")
				c.connected <- true
			case *slack.MessageEvent:
				c.handleMessage(elastabot.SlackCommandParser{}, ev)
			default:
			}
		}
	}
}

func (c *Client) handleMessage(parser elastabot.CommandParser, event *slack.MessageEvent) {
	command, err := parser.Parse(event.Msg.Text)
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

func (c *Client) respond(message string) {
	c.rtm.SendMessage(c.rtm.NewOutgoingMessage(message, c.channel))
}

// Stop shuts down the client
func (c *Client) Stop() {
	log.Info("shutdown signal received")
	c.shutdown <- true
}
