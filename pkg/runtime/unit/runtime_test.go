package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/runtime"
	"github.com/funtimecoding/soil/pkg/runtime/unit/runtime_tester"
	"testing"
)

func TestFunctionName(t *testing.T) {
	assert.String(
		t,
		"github.com/funtimecoding/soil/pkg/runtime/unit/runtime_tester.Fixture",
		runtime.FunctionName(runtime_tester.Fixture),
	)
}

func TestRunningFromBinary(t *testing.T) {
	assert.False(t, runtime.RunningFromBinary())
}
