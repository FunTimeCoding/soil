package network

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestSplitHostPort(t *testing.T) {
	host, port := SplitHostPort("127.0.0.1:80")
	assert.String(t, "127.0.0.1", host)
	assert.Integer(t, 80, port)
}
