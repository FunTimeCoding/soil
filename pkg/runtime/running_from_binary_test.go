package runtime

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func TestRunningFromBinary(t *testing.T) {
	assert.False(t, RunningFromBinary())
}
