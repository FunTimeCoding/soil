package network

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestSplitHost(t *testing.T) {
	assert.String(t, "127.0.0.1", SplitHost("127.0.0.1:80"))
}
