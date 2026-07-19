package network

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestSplitPort(t *testing.T) {
	assert.Integer(t, 80, SplitPort("127.0.0.1:80"))
}
