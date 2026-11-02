package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/log/logger"
	usage "github.com/funtimecoding/soil/pkg/nextcloud/usage/mock_client"
	prometheus "github.com/funtimecoding/soil/pkg/prometheus/mock_client"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/daemon"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/unit/service_tester"
	"go.uber.org/goleak"
	"testing"
)

func TestTheDaemonClosesWhatItBuiltOnce(t *testing.T) {
	p := newPrometheus(t)
	u := newUsage(t)
	d := daemon.New(
		service_tester.NewOption(t, service_tester.FixtureBoard(), p, nil, u),
		logger.New(t.Context()),
		memory.New(),
	)
	d.Close()
	d.Close()
	assert.Integer(t, 1, p.Closed)
	assert.Integer(t, 1, u.Closed)
}

func TestTheDaemonLeavesNothingRunning(t *testing.T) {
	before := goleak.IgnoreCurrent()
	p := prometheus.New()
	u := usage.New()
	d := daemon.New(
		service_tester.NewOption(t, service_tester.FixtureBoard(), p, nil, u),
		logger.New(t.Context()),
		memory.New(),
	)
	workers := d.Workers()

	for _, w := range workers {
		w.Start()
	}

	for _, w := range workers {
		w.Close()
	}

	d.Close()
	p.Close()
	u.Close()
	goleak.VerifyNone(t, before)
}

func TestTheDaemonHandsOutItsGuardsThenItsClock(t *testing.T) {
	d := daemon.New(
		service_tester.NewOption(
			t,
			service_tester.FixtureBoard(),
			newPrometheus(t),
			nil,
			newUsage(t),
		),
		logger.New(t.Context()),
		memory.New(),
	)
	t.Cleanup(d.Close)
	assert.Count(t, 3, d.Workers())
}
