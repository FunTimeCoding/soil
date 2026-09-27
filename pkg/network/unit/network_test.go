package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/network"
	"testing"
)

func TestSplitHost(t *testing.T) {
	assert.String(t, "127.0.0.1", network.SplitHost("127.0.0.1:80"))
}

func TestSplitPort(t *testing.T) {
	assert.Integer(t, 80, network.SplitPort("127.0.0.1:80"))
}

func TestSplitHostPort(t *testing.T) {
	host, port := network.SplitHostPort("127.0.0.1:80")
	assert.String(t, "127.0.0.1", host)
	assert.Integer(t, 80, port)
}
