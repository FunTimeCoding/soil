package socket

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"strings"
	"testing"
)

func TestPathDeterministic(t *testing.T) {
	a := Path("/tmp/project/Procfile")
	b := Path("/tmp/project/Procfile")
	assert.String(t, a, b)
}

func TestPathDifferentDirectories(t *testing.T) {
	a := Path("/tmp/alpha/Procfile")
	b := Path("/tmp/bravo/Procfile")
	assert.True(t, a != b)
}

func TestPathEndsWith(t *testing.T) {
	result := Path("/tmp/project/Procfile")
	assert.True(t, strings.HasSuffix(result, ".sock"))
}
