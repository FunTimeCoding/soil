package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/socket"
	"strings"
	"testing"
)

func TestPathDeterministic(t *testing.T) {
	a := socket.Path("/tmp/project/Procfile")
	b := socket.Path("/tmp/project/Procfile")
	assert.String(t, a, b)
}

func TestPathDifferentDirectories(t *testing.T) {
	a := socket.Path("/tmp/alpha/Procfile")
	b := socket.Path("/tmp/bravo/Procfile")
	assert.True(t, a != b)
}

func TestPathEndsWith(t *testing.T) {
	result := socket.Path("/tmp/project/Procfile")
	assert.True(t, strings.HasSuffix(result, ".sock"))
}
