package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/runtime"
	"testing"
)

func Fixture() {}

func TestFunctionName(t *testing.T) {
	assert.String(
		t,
		"github.com/funtimecoding/soil/pkg/runtime/unit_test_test.Fixture",
		runtime.FunctionName(Fixture),
	)
}
