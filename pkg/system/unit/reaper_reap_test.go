package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/system/reaper"
	"github.com/funtimecoding/soil/pkg/system/run"
	"github.com/funtimecoding/soil/pkg/system/types/zombie_detail"
	"github.com/prometheus/client_golang/prometheus"
	"testing"
	"time"
)

func TestReapedChildMovesCounterAndCapturesNothing(t *testing.T) {
	windowsSkip(t)
	child := run.New().Open("true").Pid()
	m := memory.New()
	y := prometheus.NewRegistry()
	r := reaper.New(m, y).WithZombies(
		func() map[int]zombie_detail.Detail {
			return map[int]zombie_detail.Detail{child: {Comm: "true"}}
		},
	)
	defer r.Close()
	deadline := time.Now().Add(5 * time.Second)

	for reaped(t, y, "true") == 0 && time.Now().Before(deadline) {
		r.Reap()
		time.Sleep(10 * time.Millisecond)
	}

	assert.Integer(t, 1, reaped(t, y, "true"))
	assert.Count(t, 0, m.Events())
}
