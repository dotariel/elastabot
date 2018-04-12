package elastabot

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConvertToMinutes(t *testing.T) {
	testCases := []struct {
		input    string
		expected int
	}{
		{input: "", expected: DefaultDuration},
		{input: "foo", expected: DefaultDuration},
		{input: "xd", expected: DefaultDuration},
		{input: "10", expected: 10},
		{input: "10m", expected: 10},
		{input: "10h", expected: 10 * 60},
		{input: "10d", expected: 10 * 60 * 24},
		{input: "10w", expected: 10 * 60 * 24 * 7},
	}

	for _, tt := range testCases {
		assert.Equal(t, tt.expected, convertToMinutes(tt.input))
	}
}
