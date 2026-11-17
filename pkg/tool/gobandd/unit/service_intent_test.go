package unit

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/assert"
	bandConstant "github.com/funtimecoding/soil/pkg/band/constant"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/unit/service_tester"
	"testing"
)

func TestShutdownRecordsHostOffline(t *testing.T) {
	o := service_tester.Intent(t)
	v := o.Service()
	_, e := v.Power("alfa", constant.OperationShutdown, false)
	assert.FatalOnError(t, e)
	assert.String(t, "offline", *o.Netbox.Updated["alfa"].Status)
}

func TestStartRecordsHostActive(t *testing.T) {
	o := service_tester.Intent(t)
	v := o.Service()
	o.Alfa().SetPower(bandConstant.PowerOffSoft, bandConstant.PowerOn)
	_, e := v.Power("alfa", constant.OperationStart, false)
	assert.FatalOnError(t, e)
	assert.String(t, "active", *o.Netbox.Updated["alfa"].Status)
}

func TestRefusedPowerRequestRecordsNothing(t *testing.T) {
	o := service_tester.Intent(t)
	v := o.Service()
	o.Alfa().SetPower(bandConstant.PowerOn, service_tester.HardStates()...)
	_, e := v.Power("alfa", constant.OperationShutdown, false)
	assert.NotNil(t, e)
	assert.Count(t, 0, o.Netbox.Updated)
}

func TestRebootRecordsNothing(t *testing.T) {
	o := service_tester.Intent(t)
	v := o.Service()
	_, e := v.Power("alfa", constant.OperationReboot, false)
	assert.FatalOnError(t, e)
	assert.Count(t, 0, o.Netbox.Updated)
}

func TestPowerActionWithoutSwitchRecordsNothing(t *testing.T) {
	o := service_tester.New(t)
	v := o.Service()
	o.Alfa().SetPower(bandConstant.PowerOn, service_tester.RunningStates()...)
	_, e := v.Power("alfa", constant.OperationShutdown, false)
	assert.FatalOnError(t, e)
	assert.Integers(t, []int{12}, o.Alfa().Requested())
}

func TestFailedRecordIsReportedAndNeverHoldsPowerBack(t *testing.T) {
	o := service_tester.Intent(t)
	v := o.Service()
	o.Netbox.Failure = errors.New("netbox answered nonsense")
	_, e := v.Power("alfa", constant.OperationShutdown, false)
	assert.FatalOnError(t, e)
	assert.Integers(t, []int{12}, o.Alfa().Requested())
	assert.Count(t, 1, o.Reporter.Events())
}

func TestUnreachableNetboxIsNotReported(t *testing.T) {
	o := service_tester.Intent(t)
	v := o.Service()
	o.Netbox.Failure = unavailable.Format("netbox restarting")
	_, e := v.Power("alfa", constant.OperationShutdown, false)
	assert.FatalOnError(t, e)
	assert.Integers(t, []int{12}, o.Alfa().Requested())
	assert.Count(t, 0, o.Reporter.Events())
}
