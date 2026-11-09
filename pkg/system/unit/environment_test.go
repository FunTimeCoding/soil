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

func TestUnsetBooleanTakesFallback(t *testing.T) {
	assert.True(t, environment.FallbackBoolean("DOES_NOT_EXIST", true))
	assert.False(t, environment.FallbackBoolean("DOES_NOT_EXIST", false))
}

func TestSetBooleanOverridesFallback(t *testing.T) {
	t.Setenv("SOIL_FALLBACK_BOOLEAN", "true")
	assert.True(t, environment.FallbackBoolean("SOIL_FALLBACK_BOOLEAN", false))
	t.Setenv("SOIL_FALLBACK_BOOLEAN", "false")
	assert.False(t, environment.FallbackBoolean("SOIL_FALLBACK_BOOLEAN", true))
}
