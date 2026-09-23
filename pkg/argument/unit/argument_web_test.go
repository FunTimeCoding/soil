package unit

import (
	argumentConstant "github.com/funtimecoding/soil/pkg/argument/constant"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"testing"
)

func TestWebDefaults(t *testing.T) {
	t.Setenv(constant.PortEnvironment, "")
	t.Setenv(constant.BindEnvironment, "")
	a := testInstance(t)
	a.Web()
	assert.Nil(t, a.ParseArguments(nil))
	assert.Integer(t, 8080, a.GetInteger(argumentConstant.Port))
	assert.String(t, "127.0.0.1", a.GetString(argumentConstant.BindAddress))
	assert.String(t, "127.0.0.1:8080", a.Address())
}

func TestWebEnvironmentOverridesDefault(t *testing.T) {
	t.Setenv(constant.PortEnvironment, "9000")
	t.Setenv(constant.BindEnvironment, "0.0.0.0")
	a := testInstance(t)
	a.Web()
	assert.Nil(t, a.ParseArguments(nil))
	assert.Integer(t, 9000, a.GetInteger(argumentConstant.Port))
	assert.String(t, "0.0.0.0", a.GetString(argumentConstant.BindAddress))
	assert.String(t, "0.0.0.0:9000", a.Address())
}

func TestWebFlagOverridesEnvironment(t *testing.T) {
	t.Setenv(constant.PortEnvironment, "9000")
	t.Setenv(constant.BindEnvironment, "0.0.0.0")
	a := testInstance(t)
	a.Web()
	assert.Nil(
		t,
		a.ParseArguments(
			[]string{"--port", "7000", "--bind-address", "192.168.0.1"},
		),
	)
	assert.Integer(t, 7000, a.GetInteger(argumentConstant.Port))
	assert.String(t, "192.168.0.1", a.GetString(argumentConstant.BindAddress))
	assert.String(t, "192.168.0.1:7000", a.Address())
}
