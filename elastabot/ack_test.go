package elastabot

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAck_Equals(t *testing.T) {
	testCases := []struct {
		input    Ack
		other    Command
		expected bool
	}{
		{input: Ack{}, other: Ack{}, expected: true},
		{input: Ack{Alert: ""}, other: Ack{}, expected: true},
		{input: Ack{Alert: "", Duration: 0}, other: Ack{}, expected: true},
		{input: Ack{Alert: "", Duration: 0, Triage: nil}, other: Ack{}, expected: true},
		{input: Ack{Alert: "Foo", Duration: 0, Triage: nil}, other: Ack{Alert: "Foo"}, expected: true},
		{input: Ack{Alert: "Foo", Duration: 10, Triage: nil}, other: Ack{Alert: "Foo", Duration: 10}, expected: true},
		{input: Ack{}, other: Triage{}, expected: false},
		{input: Ack{Alert: "Foo"}, other: Ack{Alert: "Bar"}, expected: false},
		// {input: Ack{Alert: "Foo", Duration: 0}, other: Ack{Alert: "Foo", Duration: 1}, expected: false},
	}

	for _, tt := range testCases {
		assert.Equal(t, tt.expected, tt.input.Equals(tt.other))
	}
}

func TestConvertToMinutes(t *testing.T) {
	testCases := []struct {
		input    string
		expected int
	}{
		{input: "foo", expected: 0},
	}

	for _, tt := range testCases {
		assert.Equal(t, tt.expected, convertToMinutes(tt.input))
	}
}
