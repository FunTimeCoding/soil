package runner_tester

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/system/reaper"
	"github.com/funtimecoding/soil/pkg/system/run"
	"github.com/prometheus/client_golang/prometheus"
	"testing"
)

func New(
	t *testing.T,
	o ...Option,
) *Tester {
	t.Helper()
	s := &setup{hold: func() {}}

	for _, f := range o {
		f(s)
	}

	if s.sleeper {
		r := reaper.New(memory.New(), prometheus.NewRegistry())
		t.Cleanup(r.Close)
		t.Cleanup(r.Interrupt)
		s.hold = func() {
			run.New().NoPanic().WithRegistry(r).Start("sleep", "30")
		}
		s.registry = r
	}

	result := build(t, s.hold, s.registry, s.downstream, s.copyModes)

	if s.gate != nil {
		result.gate = s.gate
		t.Cleanup(result.Release)
	}

	return result
}
