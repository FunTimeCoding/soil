package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/runtime"
	"testing"
)

func Fixture() {}

func TestFunctionName(t *testing.T) {
	assert.String(
		t,
		"github.com/funtimecoding/soil/pkg/runtime/unit.Fixture",
		runtime.FunctionName(Fixture),
	)
}
