package system

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestLookPath(t *testing.T) {
	assert.True(t, len(LookPath("ls")) > 0)
}
