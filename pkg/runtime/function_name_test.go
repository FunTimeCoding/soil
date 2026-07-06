package runtime

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"testing"
)

func Fixture() {}

func TestFunctionName(t *testing.T) {
	assert.String(
		t,
		"github.com/funtimecoding/soil/pkg/runtime.Fixture",
		FunctionName(Fixture),
	)
}
