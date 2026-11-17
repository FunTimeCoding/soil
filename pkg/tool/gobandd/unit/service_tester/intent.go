package service_tester

import (
	bandConstant "github.com/funtimecoding/soil/pkg/band/constant"
	"testing"
)

func Intent(t *testing.T) *Tester {
	t.Helper()
	result := New(t, WithNetbox())
	result.Alfa().SetPower(bandConstant.PowerOn, RunningStates()...)

	return result
}
