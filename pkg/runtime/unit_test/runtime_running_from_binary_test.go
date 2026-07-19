package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/runtime"
	"testing"
)

func TestRunningFromBinary(t *testing.T) {
	assert.False(t, runtime.RunningFromBinary())
}
