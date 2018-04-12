package elastabot

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockHandler struct {
	subject string
}

func (h *mockHandler) StartTriage(s string) error {
	h.subject = s
	return nil
}

func TestTriage_Equals(t *testing.T) {
	testCases := []struct {
		input    Triage
		other    Command
		expected bool
	}{
		{input: Triage{}, other: Triage{}, expected: true},
		{input: Triage{Topic: "foo"}, other: Triage{Topic: "foo"}, expected: true},
		{input: Triage{Topic: "foo"}, other: Triage{Topic: "bar"}, expected: false},
	}

	for _, tt := range testCases {
		assert.Equal(t, tt.expected, tt.input.Equals(tt.other))
	}
}
func TestExecute(t *testing.T) {
	handler := &mockHandler{}
	triage := NewTriage("foo", handler)

	result, err := triage.Execute()

	assert.Nil(t, err)
	assert.Equal(t, "ok", result)
	assert.Equal(t, "foo", handler.subject)
}
