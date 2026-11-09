package unit

import (
	"github.com/funtimecoding/soil/pkg/argument"
	argumentConstant "github.com/funtimecoding/soil/pkg/argument/constant"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/constant"
	"testing"
)

func TestBoardComesFromEnvironment(t *testing.T) {
	t.Setenv(
		constant.BoardPathEnvironment,
		"/etc/godashboard/godashboardd.yaml",
	)
	a := argument.NewInstance(constant.Identity)
	godashboardd.Arguments(a)
	assert.Nil(t, a.ParseArguments(nil))
	assert.String(
		t,
		"/etc/godashboard/godashboardd.yaml",
		a.GetString(argumentConstant.Board),
	)
}

func TestBoardFlagOverridesEnvironment(t *testing.T) {
	t.Setenv(
		constant.BoardPathEnvironment,
		"/etc/godashboard/godashboardd.yaml",
	)
	a := argument.NewInstance(constant.Identity)
	godashboardd.Arguments(a)
	assert.Nil(t, a.ParseArguments([]string{"--board", "/explicit.yaml"}))
	assert.String(t, "/explicit.yaml", a.GetString(argumentConstant.Board))
}
