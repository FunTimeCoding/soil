package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/strings/upper"
	"github.com/funtimecoding/soil/pkg/system/environment"
	"testing"
)

func TestEnvironment(t *testing.T) {
	assert.String(
		t,
		"Alfa",
		environment.Fallback("DOES_NOT_EXIST", upper.Alfa),
	)
	environment.EnsureUnset("NEVER_EXIST")
}
