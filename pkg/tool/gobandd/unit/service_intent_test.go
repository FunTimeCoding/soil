package unit

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/assert"
	band "github.com/funtimecoding/soil/pkg/band/constant"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/service"
	"testing"
)

func TestAShutdownRecordsTheHostOffline(t *testing.T) {
	v, n, _, _ := intentService(t)
	_, e := v.Power("alfa", constant.OperationShutdown, false)
	assert.FatalOnError(t, e)
	assert.String(t, "offline", *n.Updated["alfa"].Status)
}

func TestAStartRecordsTheHostActive(t *testing.T) {
	v, n, engine, _ := intentService(t)
	engine.SetPower(band.PowerOffSoft, band.PowerOn)
	_, e := v.Power("alfa", constant.OperationStart, false)
	assert.FatalOnError(t, e)
	assert.String(t, "active", *n.Updated["alfa"].Status)
}

func TestARefusedPowerRequestRecordsNothing(t *testing.T) {
	v, n, engine, _ := intentService(t)
	engine.SetPower(band.PowerOn, hardStates()...)
	_, e := v.Power("alfa", constant.OperationShutdown, false)
	assert.NotNil(t, e)
	assert.Count(t, 0, n.Updated)
}

func TestARebootRecordsNothing(t *testing.T) {
	v, n, _, _ := intentService(t)
	_, e := v.Power("alfa", constant.OperationReboot, false)
	assert.FatalOnError(t, e)
	assert.Count(t, 0, n.Updated)
}

func TestAPowerActionWithoutTheSwitchRecordsNothing(t *testing.T) {
	engine := mock_client.New()
	engine.SetPower(band.PowerOn, runningStates()...)
	v := service.New(alfaOption(), logger.New(t.Context()), memory.New())
	defer v.Close()
	v.Pool().Set("alfa", engine)
	_, e := v.Power("alfa", constant.OperationShutdown, false)
	assert.FatalOnError(t, e)
	assert.Count(t, 0, v.Workers())
}

func TestAFailedRecordIsReportedAndNeverHoldsPowerBack(t *testing.T) {
	v, n, engine, r := intentService(t)
	n.Failure = errors.New("netbox answered nonsense")
	_, e := v.Power("alfa", constant.OperationShutdown, false)
	assert.FatalOnError(t, e)
	assert.Integers(t, []int{12}, engine.Requested())
	assert.Count(t, 1, r.Events())
}

func TestAnUnreachableNetboxIsNotReported(t *testing.T) {
	v, n, engine, r := intentService(t)
	n.Failure = unavailable.Format("netbox restarting")
	_, e := v.Power("alfa", constant.OperationShutdown, false)
	assert.FatalOnError(t, e)
	assert.Integers(t, []int{12}, engine.Requested())
	assert.Count(t, 0, r.Events())
}
