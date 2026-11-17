package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/event/notifier"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert"
	prometheusConstant "github.com/funtimecoding/soil/pkg/prometheus/constant"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/unit/service_tester"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/worker"
	"go.uber.org/goleak"
	"testing"
	"time"
)

func TestDaemonClosesWhatItBuilt(t *testing.T) {
	s := service_tester.New(t)
	d := s.Daemon()
	d.Close()
	d.Close()
	assert.Integer(t, 1, s.Alertmanager.Closed)
}

func TestDaemonHandsOutItsGuardThenItsClock(t *testing.T) {
	workers := service_tester.New(t).Daemon().Workers()
	assert.Count(t, 2, workers)
	_, guard := workers[0].(*peer.Guard)
	assert.True(t, guard)
	_, clock := workers[1].(*worker.Worker)
	assert.True(t, clock)
}

func TestDaemonLeavesNothingRunning(t *testing.T) {
	s := service_tester.New(t)
	before := goleak.IgnoreCurrent()
	d := s.Daemon()
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

func TestRecordStampsLastPollAndNotifies(t *testing.T) {
	n := notifier.New()
	changes := n.Subscribe()
	defer n.Unsubscribe(changes)
	now := time.Date(2026, time.November, 2, 6, 0, 0, 0, time.UTC)
	v := service_tester.New(
		t,
		service_tester.WithClock(func() time.Time { return now }),
		service_tester.WithNotifier(n),
	).Service()
	assert.True(t, v.LastPoll().IsZero())
	_, e := v.Record()
	assert.FatalOnError(t, e)
	assert.True(t, v.LastPoll().Equal(now))
	assert.Integer(t, 1, len(changes))
}

func TestFailedRecordStillStampsAndNotifies(t *testing.T) {
	n := notifier.New()
	changes := n.Subscribe()
	defer n.Unsubscribe(changes)
	s := service_tester.New(t, service_tester.WithNotifier(n))
	s.Alertmanager.Fail(errors.Format("alertmanager refused", "test"))
	v := s.Service()
	_, e := v.Record()
	assert.NotNil(t, e)
	assert.False(t, v.LastPoll().IsZero())
	assert.Integer(t, 1, len(changes))
}

func TestRecordReportsWhatItChanged(t *testing.T) {
	s := service_tester.New(t)
	s.Alertmanager.Add(
		alert.NewBasic(
			"abc123",
			"HighMemory",
			prometheusConstant.CriticalSeverity,
			"",
		),
	)
	s.Alertmanager.Add(
		alert.NewBasic(
			"def456",
			"DiskFull",
			prometheusConstant.WarningSeverity,
			"",
		),
	)
	v := s.Service()
	o, e := v.Record()
	assert.FatalOnError(t, e)
	assert.Integer(t, 2, o.Created)
	assert.Integer(t, 2, o.Firing)
	assert.Integer(t, 2, o.Total)
	s.Alertmanager.Remove("abc123")
	p, f := v.Record()
	assert.FatalOnError(t, f)
	assert.Integer(t, 0, p.Created)
	assert.Integer(t, 1, p.Firing)
	assert.Integer(t, 2, p.Total)
}

func TestFailedAlertmanagerReadIsReturned(t *testing.T) {
	s := service_tester.New(t)
	s.Alertmanager.Fail(errors.Format("alertmanager refused", "test"))
	_, e := s.Service().Record()
	assert.StringContains(t, "alertmanager refused", e.Error())
}
