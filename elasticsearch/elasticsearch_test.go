package elasticsearch

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConnect(t *testing.T) {
	url := "http://any.ip:9200"
	conn, err := Connect(url)

	assert.Nil(t, err)
	assert.NotNil(t, conn)
}

func TestPing(t *testing.T) {
	conn := CannedResponder("ping")
	assert.True(t, conn.Ping())
}
