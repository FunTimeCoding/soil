package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/peer/unit/guard_tester"
	"github.com/funtimecoding/soil/pkg/peer/unit/switchable_peer"
	"github.com/funtimecoding/soil/pkg/time/clock_tester"
	"github.com/funtimecoding/soil/pkg/web/outage_tester"
	"sync/atomic"
	"testing"
	"time"
)

func TestUpstreamUnavailableMarksGuardDown(t *testing.T) {
	r := memory.New()
	g := peer.New(
		"kilo",
		"http://server.test:8080",
		switchable_peer.New(),
		logger.New(context.Background()),
		r,
	)
	defer g.Close()
	calls := 0
	call := func(context.Context) error {
		calls++

		return outage_tester.New(t).Unavailable()
	}
	e := g.Attempt(t.Context(), call)
	assert.True(t, unavailable.Is(e))
	assert.StringContains(t, "kilo unreachable since", e.Error())
	assert.True(t, unavailable.Is(g.Attempt(t.Context(), call)))
	assert.Integer(t, 1, calls)
	assert.Count(t, 0, r.Events())
}

func TestProbeDoesNotBringUnavailablePeerBack(t *testing.T) {
	p := switchable_peer.New()
	g := peer.New(
		"oscar",
		"http://server.test:8080",
		p,
		logger.New(context.Background()),
		memory.New(),
	).WithBackoff(time.Millisecond, 4*time.Millisecond)
	g.Start()
	defer g.Close()
	assert.True(
		t,
		unavailable.Is(
			g.Attempt(
				t.Context(),
				func(context.Context) error {
					return outage_tester.New(t).Unavailable()
				},
			),
		),
	)
	time.Sleep(50 * time.Millisecond)
	assert.True(t, unavailable.Is(g.Unavailable()))
	assert.Integer(t, 0, p.Probes())
}

func TestNextSuccessfulCallBringsUnavailablePeerBack(t *testing.T) {
	c := clock_tester.New()
	p := switchable_peer.New()
	g := peer.New(
		"lima",
		"http://server.test:8080",
		p,
		logger.New(context.Background()),
		memory.New(),
	).WithBackoff(time.Second, 8*time.Second).WithClock(c.Now)
	g.Start()
	defer g.Close()
	failing := true
	calls := 0
	call := func(context.Context) error {
		calls++

		if failing {
			return outage_tester.New(t).Unavailable()
		}

		return nil
	}
	assert.True(t, unavailable.Is(g.Attempt(t.Context(), call)))
	c.Advance(500 * time.Millisecond)
	assert.True(t, unavailable.Is(g.Attempt(t.Context(), call)))
	assert.Integer(t, 1, calls)
	c.Advance(time.Second)
	assert.True(t, unavailable.Is(g.Attempt(t.Context(), call)))
	assert.Integer(t, 2, calls)
	c.Advance(time.Second)
	assert.True(t, unavailable.Is(g.Attempt(t.Context(), call)))
	assert.Integer(t, 2, calls)
	c.Advance(time.Second)
	failing = false
	assert.Nil(t, g.Attempt(t.Context(), call))
	assert.Integer(t, 3, calls)
	assert.Nil(t, g.Unavailable())
	assert.Integer(t, 0, p.Probes())
}

func TestPanickingTrialLeavesNextTrialAdmitted(t *testing.T) {
	c := clock_tester.New()
	g := peer.New(
		"papa",
		"http://server.test:8080",
		switchable_peer.New(),
		logger.New(context.Background()),
		memory.New(),
	).WithBackoff(time.Second, time.Minute).WithClock(c.Now)
	defer g.Close()
	assert.True(
		t,
		unavailable.Is(
			g.Attempt(
				t.Context(),
				func(context.Context) error {
					return outage_tester.New(t).Unavailable()
				},
			),
		),
	)
	c.Advance(time.Second)
	assert.Any(t, "trial", guard_tester.Panicked(t.Context(), g))
	calls := 0
	assert.Nil(
		t,
		g.Attempt(
			t.Context(),
			func(context.Context) error {
				calls++

				return nil
			},
		),
	)
	assert.Integer(t, 1, calls)
	assert.Nil(t, g.Unavailable())
}

func TestLastingUnavailablePeerIsCapturedOnce(t *testing.T) {
	c := clock_tester.New()
	r := memory.New()
	g := peer.New(
		"mike",
		"http://server.test:8080",
		switchable_peer.New(),
		logger.New(context.Background()),
		r,
	).WithBackoff(time.Second, time.Minute).WithClock(c.Now)
	defer g.Close()
	call := func(context.Context) error { return outage_tester.New(t).Unavailable() }

	for range 20 {
		assert.True(t, unavailable.Is(g.Attempt(t.Context(), call)))
		c.Advance(time.Minute)
	}

	assert.Count(t, 1, r.Events())
	assert.StringContains(t, "mike down for 5m0s", r.Events()[0].Error.Error())
}

func TestOnlyOneTrialReachesUnavailablePeer(t *testing.T) {
	c := clock_tester.New()
	g := peer.New(
		"november",
		"http://server.test:8080",
		switchable_peer.New(),
		logger.New(context.Background()),
		memory.New(),
	).WithBackoff(time.Second, time.Minute).WithClock(c.Now)
	defer g.Close()
	assert.True(
		t,
		unavailable.Is(
			g.Attempt(
				t.Context(),
				func(context.Context) error {
					return outage_tester.New(t).Unavailable()
				},
			),
		),
	)
	c.Advance(time.Second)
	var calls atomic.Int32
	entered := make(chan struct{})
	release := make(chan struct{})
	done := make(chan error)
	go func() {
		done <- g.Attempt(
			t.Context(),
			func(context.Context) error {
				calls.Add(1)
				close(entered)
				<-release

				return nil
			},
		)
	}()
	<-entered
	assert.True(
		t,
		unavailable.Is(
			g.Attempt(
				t.Context(),
				func(context.Context) error {
					calls.Add(1)

					return nil
				},
			),
		),
	)
	close(release)
	assert.Nil(t, <-done)
	assert.Integer(t, 1, int(calls.Load()))
	assert.Nil(t, g.Unavailable())
}
