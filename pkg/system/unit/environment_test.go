package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	stringsConstant "github.com/funtimecoding/soil/pkg/strings/constant"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"testing"
)

func TestEnvironment(t *testing.T) {
	assert.String(
		t,
		"Alfa",
		environment.Fallback("DOES_NOT_EXIST", stringsConstant.UpperAlfa),
	)
	environment.EnsureUnset("NEVER_EXIST")
}
