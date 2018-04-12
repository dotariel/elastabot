package slack

import (
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/gorilla/websocket"
	"github.com/nlopes/slack"
	log "github.com/sirupsen/logrus"
)

type MockServer struct {
	*httptest.Server
	responses chan string
}

func newMockServer() *MockServer {
	server := &MockServer{
		Server:    httptest.NewServer(nil),
		responses: make(chan string, 10),
	}

	slack.SLACK_API = server.httpURL()

	http.HandleFunc("/rtm", server.websocket)
	http.HandleFunc("/rtm.connect", server.connect)

	log.Printf("mock server started; %s, %s", server.httpURL(), server.wsURL())

	return server
}

func (ms *MockServer) httpURL() string {
	return "http://" + ms.Listener.Addr().String() + "/"
}

func (ms *MockServer) wsURL() string {
	return "ws://" + ms.Listener.Addr().String() + "/rtm"
}

func (ms *MockServer) record(message string) {
	ms.responses <- message
}

func (ms *MockServer) connect(w http.ResponseWriter, r *http.Request) {
	response := fmt.Sprintf(`{ "ok": true, "url":"%v"}`, ms.wsURL())
	bytes := []byte(response)
	w.Write(bytes)
}

func (ms *MockServer) websocket(w http.ResponseWriter, r *http.Request) {
	log.Debug("handling request for", r.URL)

	upgrader := websocket.Upgrader{
		CheckOrigin: func(*http.Request) bool { return true },
	}
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
	}
	defer c.Close()

	for {
		_, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}

		// Right here is where we can respond
		ms.record(string(message))

		// TODO: Write back
	}
}

func (ms *MockServer) stop() {
	ms.Close()
}

func newMockClient() *Client {
	return &Client{
		rtm:       slack.New("token").NewRTM(),
		connected: make(chan bool),
		shutdown:  make(chan bool),
	}
}

func newMessageEvent(body string) slack.RTMEvent {
	return slack.RTMEvent{
		Type: "Message",
		Data: &slack.MessageEvent{
			Msg: slack.Msg{
				User:    "_user",
				Channel: "_channel",
				Text:    body,
			},
		},
	}
}
