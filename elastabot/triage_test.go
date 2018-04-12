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

func TestExecute(t *testing.T) {
	handler := &mockHandler{}
	triage := NewTriage("foo", handler)

	result, err := triage.Execute()

	assert.Nil(t, err)
	assert.Equal(t, "ok", result)
	assert.Equal(t, "foo", handler.subject)
}
