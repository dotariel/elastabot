package elastabot

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseCommand(t *testing.T) {
	testCases := []struct {
		message  string
		expected Command
		err      error
	}{
		{
			message:  "nomatch",
			expected: nil,
			err:      errors.New("unknown command"),
		},
		{
			message:  "ack",
			expected: nil,
			err:      errors.New("unknown command"),
		},
		{
			message: "!ack",
			expected: &Ack{
				Alert:    "",
				Duration: 0,
				Triage:   nil,
			},
			err: nil,
		},
		{
			message: "!ack Acme Flatline Alert",
			expected: &Ack{
				Alert:    "Acme Flatline Alert",
				Duration: 0,
				Triage:   nil,
			},
			err: nil,
		},
		{
			message: "!ack Acme?",
			expected: &Ack{
				Alert:    "Acme",
				Duration: 0,
				Triage: &Triage{
					Topic:   "Acme",
					Handler: nil,
				},
			},
			err: nil,
		},
		{
			message: "!ack Acme|10",
			expected: &Ack{
				Alert:    "Acme",
				Duration: 10,
				Triage:   nil,
			},
			err: nil,
		},
	}

	for _, tt := range testCases {
		resp, err := ParseCommand(tt.message)
		assert.Equal(t, tt.err, err)
		assert.Equal(t, tt.expected, resp)
	}

}
