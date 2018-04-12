package elastabot

import (
	"strconv"
	"strings"
)

const (
	DefaultDuration = 0
)

type Ack struct {
	Alert    string
	Duration int
	Triage   *Triage
}

func NewAck(alert string, duration string, triage *Triage) *Ack {
	return &Ack{
		Alert:    alert,
		Duration: convertToMinutes(duration),
		Triage:   triage,
	}
}

func (a Ack) Execute() (string, error) {
	// Talk to ES...
	// Also handle or delegate to Triage.execute()
	return "", nil
}

func convertToMinutes(d string) int {
	if len(d) == 0 {
		return DefaultDuration // TODO: Return the default value
	}

	if strings.HasSuffix(d, "m") {
		return valueOrDefault(d[:len(d)-1], 1)
	}

	if strings.HasSuffix(d, "h") {
		return valueOrDefault(d[:len(d)-1], 60)
	}

	if strings.HasSuffix(d, "d") {
		return valueOrDefault(d[:len(d)-1], 60*24)
	}

	if strings.HasSuffix(d, "w") {
		return valueOrDefault(d[:len(d)-1], 60*24*7)
	}

	return valueOrDefault(d, 1)
}

func valueOrDefault(value string, multiplier int) int {
	n, err := strconv.Atoi(value)
	if err != nil {
		return DefaultDuration
	}

	return n * multiplier
}
