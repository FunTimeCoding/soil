package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/connection"
	errorsConstant "github.com/funtimecoding/soil/pkg/errors/constant"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/daemon"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/option"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/unit/service_tester"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/unit/worker_tester"
	"go.uber.org/goleak"
	"testing"
	"time"
)

func TestTheDaemonWithoutOptionsRunsOnlyItsGuard(t *testing.T) {
	workers := service_tester.NewDaemonWithoutSubscription(
		t,
		service_tester.NewMattermost(t, "kilo"),
	).Workers()
	assert.Count(t, 1, workers)
	_, guard := workers[0].(*peer.Guard)
	assert.True(t, guard)
}

func TestGoclaudedAddsTheRosterGuardTheWatcherAndTwoClocks(t *testing.T) {
	_, k := worker_tester.NewSink(t)
	workers := service_tester.NewDaemon(
		t,
		service_tester.NewMattermost(t, "kilo"),
		k,
		nil,
	).Workers()
	assert.Count(t, 5, workers)

	for n, w := range workers {
		_, guard := w.(*peer.Guard)
		assert.Boolean(t, n < 2, guard)
	}
}

func TestTheMonitorAddsItsClock(t *testing.T) {
	_, k := worker_tester.NewSink(t)
	workers := service_tester.NewDaemon(
		t,
		service_tester.NewMattermost(t, "kilo"),
		k,
		func(o *option.Mattermost) {
			o.Monitor.Enabled = true
			o.Monitor.Interval = time.Hour
		},
	).Workers()
	assert.Count(t, 6, workers)
}

func TestTheDaemonClosesMattermostOnce(t *testing.T) {
	m := service_tester.NewMattermost(t, "kilo")
	_, k := worker_tester.NewSink(t)
	d := service_tester.NewDaemon(t, m, k, nil)
	d.Close()
	d.Close()
	assert.Integer(t, 1, m.Closed)
}

func TestTheDaemonLeavesNothingRunning(t *testing.T) {
	m := service_tester.NewMattermost(t, "kilo")
	m.Failure = connection.New(
		errorsConstant.Unreachable,
		"alfa.test",
		"",
		errorsConstant.Refused,
	)
	_, k := worker_tester.NewSink(t)
	o := service_tester.NewOption(t, m, k, nil)
	before := goleak.IgnoreCurrent()
	d := daemon.New(o, logger.New(context.Background()), memory.New())
	workers := d.Workers()

	for _, w := range workers {
		w.Start()
	}

	for _, w := range workers {
		w.Close()
	}

	d.Close()
	d.Close()
	goleak.VerifyNone(t, before)
}
