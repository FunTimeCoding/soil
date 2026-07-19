package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/network"
	"testing"
)

func TestSplitHost(t *testing.T) {
	assert.String(t, "127.0.0.1", network.SplitHost("127.0.0.1:80"))
}
