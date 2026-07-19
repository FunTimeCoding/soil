package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/system"
	"testing"
)

func TestLookPath(t *testing.T) {
	assert.True(t, len(system.LookPath("ls")) > 0)
}
