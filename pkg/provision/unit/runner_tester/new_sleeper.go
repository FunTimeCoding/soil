package runner_tester

import (
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/system/reaper"
	"github.com/funtimecoding/soil/pkg/system/run"
	"github.com/prometheus/client_golang/prometheus"
	"testing"
)

func NewSleeper(t *testing.T) *Tester {
	t.Helper()
	r := reaper.New(memory.New(), prometheus.NewRegistry())
	result := New(
		t,
		withHold(
			func() {
				run.New().NoPanic().WithRegistry(r).Start("sleep", "30")
			},
		),
		withRegistry(r),
	)
	t.Cleanup(r.Interrupt)

	return result
}
