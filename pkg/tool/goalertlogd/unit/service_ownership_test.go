package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/event/notifier"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert"
	prometheusConstant "github.com/funtimecoding/soil/pkg/prometheus/constant"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/unit/service_tester"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/worker"
	"go.uber.org/goleak"
	"testing"
	"time"
)

func TestTheDaemonClosesWhatItBuilt(t *testing.T) {
	c := service_tester.NewAlertmanager(t)
	d := service_tester.NewDaemon(
		t,
		service_tester.NewOption(service_tester.DatabasePath(t), c),
		memory.New(),
	)
	d.Close()
	d.Close()
	assert.Integer(t, 1, c.Closed)
}

func TestTheDaemonHandsOutItsGuardThenItsClock(t *testing.T) {
	workers := service_tester.NewDaemon(
		t,
		service_tester.NewOption(
			service_tester.DatabasePath(t),
			service_tester.NewAlertmanager(t),
		),
		memory.New(),
	).Workers()
	assert.Count(t, 2, workers)
	_, guard := workers[0].(*peer.Guard)
	assert.True(t, guard)
	_, clock := workers[1].(*worker.Worker)
	assert.True(t, clock)
}

func TestTheDaemonLeavesNothingRunning(t *testing.T) {
	o := service_tester.NewOption(
		service_tester.DatabasePath(t),
		service_tester.NewAlertmanager(t),
	)
	before := goleak.IgnoreCurrent()
	d := service_tester.NewDaemon(t, o, memory.New())
	workers := d.Workers()

	for _, w := range workers {
		w.Start()
	}

	for _, w := range workers {
		w.Close()
	}

	d.Close()
	goleak.VerifyNone(t, before)
}

func TestARecordStampsTheLastPollAndNotifies(t *testing.T) {
	n := notifier.New()
	changes := n.Subscribe()
	defer n.Unsubscribe(changes)
	v := service_tester.NewServiceWith(
		t,
		service_tester.DatabasePath(t),
		service_tester.NewAlertmanager(t),
		constant.Retention,
		n,
	)
	assert.True(t, v.LastPoll().IsZero())
	now := time.Date(2026, time.November, 2, 6, 0, 0, 0, time.UTC)
	_, e := v.Record(now)
	assert.FatalOnError(t, e)
	assert.True(t, v.LastPoll().Equal(now))
	assert.Integer(t, 1, len(changes))
}

func TestAFailedRecordStillStampsAndNotifies(t *testing.T) {
	c := service_tester.NewAlertmanager(t)
	c.Fail(errors.Format("alertmanager refused", "test"))
	n := notifier.New()
	changes := n.Subscribe()
	defer n.Unsubscribe(changes)
	v := service_tester.NewServiceWith(
		t,
		service_tester.DatabasePath(t),
		c,
		constant.Retention,
		n,
	)
	_, e := v.Record(time.Now())
	assert.NotNil(t, e)
	assert.False(t, v.LastPoll().IsZero())
	assert.Integer(t, 1, len(changes))
}

func TestARecordReportsWhatItChanged(t *testing.T) {
	c := service_tester.NewAlertmanager(t)
	c.Add(
		alert.NewBasic(
			"abc123",
			"HighMemory",
			prometheusConstant.CriticalSeverity,
			"",
		),
	)
	c.Add(
		alert.NewBasic(
			"def456",
			"DiskFull",
			prometheusConstant.WarningSeverity,
			"",
		),
	)
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
