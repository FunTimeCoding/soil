package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert"
	"github.com/funtimecoding/soil/pkg/prometheus/constant"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/service"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/unit/service_tester"
	"testing"
	"time"
)

func TestTheServiceClosesWhatItBuilt(t *testing.T) {
	c := service_tester.NewAlertmanager(t)
	v := service.New(
		service_tester.NewOption(service_tester.DatabasePath(t), c),
		logger.New(t.Context()),
		memory.New(),
	)
	v.Close()
	v.Close()
	assert.Integer(t, 1, c.Closed)
	_, e := v.Status()
	assert.StringContains(t, "database is closed", e.Error())
}

func TestARecordReportsWhatItChanged(t *testing.T) {
	c := service_tester.NewAlertmanager(t)
	c.Add(alert.NewBasic("abc123", "HighMemory", constant.CriticalSeverity, ""))
	c.Add(alert.NewBasic("def456", "DiskFull", constant.WarningSeverity, ""))
	v := service_tester.NewService(t, service_tester.DatabasePath(t), c)
	o, e := v.Record(time.Now())
	assert.FatalOnError(t, e)
	assert.Integer(t, 2, o.Created)
	assert.Integer(t, 2, o.Firing)
	assert.Integer(t, 2, o.Total)
	c.Remove("abc123")
	p, f := v.Record(time.Now())
	assert.FatalOnError(t, f)
	assert.Integer(t, 0, p.Created)
	assert.Integer(t, 1, p.Firing)
	assert.Integer(t, 2, p.Total)
}

func TestAFailedAlertmanagerReadIsReturned(t *testing.T) {
	c := service_tester.NewAlertmanager(t)
	c.Fail(errors.Format("alertmanager refused", "test"))
	v := service_tester.NewService(t, service_tester.DatabasePath(t), c)
	_, e := v.Record(time.Now())
	assert.StringContains(t, "alertmanager refused", e.Error())
}
