package unit_test

import (
	"github.com/funtimecoding/soil/pkg/argument"
	argumentConstant "github.com/funtimecoding/soil/pkg/argument/constant"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/identity"
	"github.com/funtimecoding/soil/pkg/web/constant"
	"testing"
)

func testInstance(t *testing.T) *argument.Instance {
	t.Helper()

	return argument.NewInstance(identity.New("gotest", "test tool", "gotest"))
}

func TestWebDefaults(t *testing.T) {
	t.Setenv(constant.PortEnvironment, "")
	t.Setenv(constant.BindEnvironment, "")
	a := testInstance(t)
	a.Web()
	assert.Nil(t, a.ParseArguments(nil))
	assert.Integer(
		t,
		constant.ListenPort,
		a.GetInteger(argumentConstant.Port),
	)
	assert.String(
		t,
		constant.Loopback,
		a.GetString(argumentConstant.BindAddress),
	)
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
