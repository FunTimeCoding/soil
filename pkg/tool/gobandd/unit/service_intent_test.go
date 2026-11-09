package unit

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/assert"
	bandConstant "github.com/funtimecoding/soil/pkg/band/constant"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/unit/service_tester"
	"testing"
)

func TestShutdownRecordsHostOffline(t *testing.T) {
	v, n, _, _ := intentService(t)
	_, e := v.Power("alfa", constant.OperationShutdown, false)
	assert.FatalOnError(t, e)
	assert.String(t, "offline", *n.Updated["alfa"].Status)
}

func TestStartRecordsHostActive(t *testing.T) {
	v, n, engine, _ := intentService(t)
	engine.SetPower(bandConstant.PowerOffSoft, bandConstant.PowerOn)
	_, e := v.Power("alfa", constant.OperationStart, false)
	assert.FatalOnError(t, e)
	assert.String(t, "active", *n.Updated["alfa"].Status)
}

func TestRefusedPowerRequestRecordsNothing(t *testing.T) {
	v, n, engine, _ := intentService(t)
	engine.SetPower(bandConstant.PowerOn, hardStates()...)
	_, e := v.Power("alfa", constant.OperationShutdown, false)
	assert.NotNil(t, e)
	assert.Count(t, 0, n.Updated)
}

func TestRebootRecordsNothing(t *testing.T) {
	v, n, _, _ := intentService(t)
	_, e := v.Power("alfa", constant.OperationReboot, false)
	assert.FatalOnError(t, e)
	assert.Count(t, 0, n.Updated)
}

func TestPowerActionWithoutSwitchRecordsNothing(t *testing.T) {
	v, engines := service_tester.New(t, memory.New(), "alfa")
	engine := engines["alfa"]
	engine.SetPower(bandConstant.PowerOn, runningStates()...)
	_, e := v.Power("alfa", constant.OperationShutdown, false)
	assert.FatalOnError(t, e)
	assert.Integers(t, []int{12}, engine.Requested())
}

func TestFailedRecordIsReportedAndNeverHoldsPowerBack(t *testing.T) {
	v, n, engine, r := intentService(t)
	n.Failure = errors.New("netbox answered nonsense")
	_, e := v.Power("alfa", constant.OperationShutdown, false)
	assert.FatalOnError(t, e)
	assert.Integers(t, []int{12}, engine.Requested())
	assert.Count(t, 1, r.Events())
}

func TestUnreachableNetboxIsNotReported(t *testing.T) {
	v, n, engine, r := intentService(t)
	n.Failure = unavailable.Format("netbox restarting")
	_, e := v.Power("alfa", constant.OperationShutdown, false)
	assert.FatalOnError(t, e)
	assert.Integers(t, []int{12}, engine.Requested())
	assert.Count(t, 0, r.Events())
}
