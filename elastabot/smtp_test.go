package elastabot

import "testing"
import "github.com/stretchr/testify/assert"

func TestSmtpOptions(t *testing.T) {

	smtp := Smtp{
		Host: "foo",
		Port: 99,
	}

	assert.Equal(t, "foo", smtp.Host)
	assert.Equal(t, 99, smtp.Port)
}
