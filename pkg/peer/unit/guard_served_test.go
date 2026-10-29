package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"github.com/funtimecoding/soil/pkg/peer/unit/fake_clock"
	"sync/atomic"
	"testing"
	"time"
)

func TestAnUpstreamUnavailableMarksTheGuardDown(t *testing.T) {
	r := memory.New()
	g := peer.New(
		"kilo",
		"http://server.test:8080",
		new(SwitchablePeer),
		logger.New(context.Background()),
		r,
	)
	defer g.Close()
	calls := 0
	call := func(context.Context) error {
		calls++

		return upstreamUnavailable()
	}
	e := g.Attempt(t.Context(), call)
	assert.True(t, unavailable.Is(e))
	assert.StringContains(t, "kilo unreachable since", e.Error())
	assert.True(t, unavailable.Is(g.Attempt(t.Context(), call)))
	assert.Integer(t, 1, calls)
	assert.Count(t, 0, r.Events())
}

func TestAProbeDoesNotBringAnUnavailablePeerBack(t *testing.T) {
	p := new(SwitchablePeer)
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
				func(context.Context) error { return upstreamUnavailable() },
			),
		),
	)
	time.Sleep(50 * time.Millisecond)
	assert.True(t, unavailable.Is(g.Unavailable()))
	assert.Integer(t, 0, p.Probes())
}

func TestTheNextSuccessfulCallBringsAnUnavailablePeerBack(t *testing.T) {
	c := fake_clock.New()
	p := new(SwitchablePeer)
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
			return upstreamUnavailable()
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

func TestALastingUnavailablePeerIsCapturedOnce(t *testing.T) {
	c := fake_clock.New()
	r := memory.New()
	g := peer.New(
		"mike",
		"http://server.test:8080",
		new(SwitchablePeer),
		logger.New(context.Background()),
		r,
	).WithBackoff(time.Second, time.Minute).WithClock(c.Now)
	defer g.Close()
	call := func(context.Context) error { return upstreamUnavailable() }

	for range 20 {
		assert.True(t, unavailable.Is(g.Attempt(t.Context(), call)))
		c.Advance(time.Minute)
	}

	assert.Count(t, 1, r.Events())
	assert.StringContains(t, "mike down for 5m0s", r.Events()[0].Error.Error())
}

func TestOnlyOneTrialReachesAnUnavailablePeer(t *testing.T) {
	c := fake_clock.New()
	g := peer.New(
		"november",
		"http://server.test:8080",
		new(SwitchablePeer),
		logger.New(context.Background()),
		memory.New(),
	).WithBackoff(time.Second, time.Minute).WithClock(c.Now)
	defer g.Close()
	assert.True(
		t,
		unavailable.Is(
			g.Attempt(
				t.Context(),
				func(context.Context) error { return upstreamUnavailable() },
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
