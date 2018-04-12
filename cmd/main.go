package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/dotariel/elastabot/slack"
)

func main() {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGINT)

	slack := slack.New()
	go slack.Start()

	<-stop
}
