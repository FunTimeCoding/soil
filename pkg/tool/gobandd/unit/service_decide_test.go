package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	band "github.com/funtimecoding/soil/pkg/band/constant"
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/service"
	"testing"
)

func TestShutdownAsksTheOperatingSystemByDefault(t *testing.T) {
	state, e := service.Decide(
		constant.OperationShutdown,
		false,
		runningStates(),
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, 12, state)
}

func TestShutdownWithForceCutsPower(t *testing.T) {
	state, e := service.Decide(
		constant.OperationShutdown,
		true,
		runningStates(),
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, 8, state)
}

func TestRebootAsksTheOperatingSystemByDefault(t *testing.T) {
	state, e := service.Decide(constant.OperationReboot, false, runningStates())
	assert.FatalOnError(t, e)
	assert.Integer(t, 14, state)
}

func TestRebootWithForcePowerCycles(t *testing.T) {
	state, e := service.Decide(constant.OperationReboot, true, runningStates())
	assert.FatalOnError(t, e)
	assert.Integer(t, 5, state)
}

func TestGracefulShutdownNeverFallsBackToACut(t *testing.T) {
	_, e := service.Decide(constant.OperationShutdown, false, hardStates())
	assert.True(t, conflict.Is(e))
	assert.StringContains(t, "force cuts power instead", e.Error())
}

func TestGracefulRebootNeverFallsBackToACycle(t *testing.T) {
	_, e := service.Decide(constant.OperationReboot, false, hardStates())
	assert.True(t, conflict.Is(e))
}

func TestStartIsRefusedWhileTheMachineRuns(t *testing.T) {
	_, e := service.Decide(constant.OperationStart, false, runningStates())
	assert.True(t, conflict.Is(e))
}

func TestStartPowersOnAMachineThatIsOff(t *testing.T) {
	state, e := service.Decide(
		constant.OperationStart,
		false,
		[]int{band.PowerOn},
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, 2, state)
}

func TestAnUnknownOperationIsTheCallersMistake(t *testing.T) {
	_, e := service.Decide("hibernate", false, runningStates())
	assert.True(t, validation.Is(e))
}
