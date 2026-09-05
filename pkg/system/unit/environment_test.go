package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"testing"
)

func TestEnvironment(t *testing.T) {
	assert.String(
		t,
		"Alfa",
		environment.Fallback("DOES_NOT_EXIST", constant.UpperAlfa),
	)
	environment.EnsureUnset("NEVER_EXIST")
}
