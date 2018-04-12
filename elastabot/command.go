package elastabot

import (
	"errors"
	"regexp"
	"strings"
)

type Command interface {
	Execute() (string, error)
}

func ParseCommand(c string) (Command, error) {
	re, _ := regexp.Compile(`^!(?P<command>ack|triage|help) ?(?P<alert>[^|?]+)? ?[|]?(?P<duration>[\d]+[smhdw]?)?`)

	keys := re.SubexpNames()
	vals := re.FindAllStringSubmatch(c, -1)

	if len(vals) == 0 {
		return nil, errors.New("unknown command")
	}

	md := map[string]string{}
	for i, n := range vals[0] {
		md[keys[i]] = n
	}

	hasTriage := strings.Contains(c, "?")

	switch md["command"] {
	case "ack":
		var triage *Triage
		var alert = md["alert"]
		var duration = md["duration"]

		if hasTriage {
			triage = &Triage{
				Topic:   alert,
				Handler: nil,
			}
		}

		return NewAck(alert, duration, triage), nil
	}
	return nil, nil
}
