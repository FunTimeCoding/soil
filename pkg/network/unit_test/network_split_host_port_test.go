package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/network"
	"testing"
)

func TestSplitHostPort(t *testing.T) {
	host, port := network.SplitHostPort("127.0.0.1:80")
	assert.String(t, "127.0.0.1", host)
	assert.Integer(t, 80, port)
}
