package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/network"
	"testing"
)

func TestSplitPort(t *testing.T) {
	assert.Integer(t, 80, network.SplitPort("127.0.0.1:80"))
}
