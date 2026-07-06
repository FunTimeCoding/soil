package environment

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings/upper"
	"testing"
)

func TestEnvironment(t *testing.T) {
	assert.String(
		t,
		"Alfa",
		Fallback("DOES_NOT_EXIST", upper.Alfa),
	)
	EnsureUnset("NEVER_EXIST")
}
