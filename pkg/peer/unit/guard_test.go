package unit

import (
	"context"
	"errors"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"testing"
	"time"
)

func TestAGuardPassesCallsThroughWhileThePeerAnswers(t *testing.T) {
	g := peer.New(
		"charlie",
		"http://server.test:8080",
		new(SwitchablePeer),
		logger.New(context.Background()),
		memory.New(),
	)
	defer g.Close()
	calls := 0
	e := g.Do(
		t.Context(),
		func(context.Context) error {
			calls++

			return nil
		},
	)
	assert.FatalOnError(t, e)
	assert.Integer(t, 1, calls)
	assert.Nil(t, g.Unavailable())
}

func TestAnOutageFailsFastWithoutReachingThePeer(t *testing.T) {
	r := memory.New()
	g := peer.New(
		"delta",
		"http://server.test:8080",
		new(SwitchablePeer),
		logger.New(context.Background()),
		r,
	)
	defer g.Close()
	calls := 0
	call := func(context.Context) error {
		calls++

		return refusedConnection()
	}
	assert.True(t, unavailable.Is(g.Do(t.Context(), call)))
	e := g.Do(t.Context(), call)
	assert.True(t, unavailable.Is(e))
	assert.StringContains(t, "delta unreachable since", e.Error())
	assert.Integer(t, 1, calls)
	assert.Count(t, 0, r.Events())
}

func TestARefusalIsReportedOnceAndThePeerStaysUp(t *testing.T) {
	r := memory.New()
	g := peer.New(
		"india",
		"http://server.test:8080",
		new(SwitchablePeer),
		logger.New(context.Background()),
		r,
	)
	defer g.Close()
	call := func(context.Context) error {
		return errors.New("model not found")
	}
	assert.StringContains(t, "model not found", g.Do(t.Context(), call).Error())
	assert.StringContains(t, "model not found", g.Do(t.Context(), call).Error())
	assert.Nil(t, g.Unavailable())
	assert.Count(t, 1, r.Events())
}

func TestAnAttemptLeavesEveryRefusalToTheCaller(t *testing.T) {
	r := memory.New()
	g := peer.New(
		"juliett",
		"http://server.test:8080",
		new(SwitchablePeer),
		logger.New(context.Background()),
		r,
	)
	defer g.Close()
	call := func(context.Context) error {
		return errors.New("content too long")
	}
	assert.StringContains(t, "too long", g.Attempt(t.Context(), call).Error())
	assert.StringContains(t, "too long", g.Attempt(t.Context(), call).Error())
	assert.Count(t, 0, r.Events())
	assert.True(
		t,
		unavailable.Is(
			g.Attempt(
				t.Context(),
				func(context.Context) error { return refusedConnection() },
			),
		),
	)
	assert.Count(t, 0, r.Events())
}

func TestAGuardProbesUntilThePeerReturns(t *testing.T) {
	p := new(SwitchablePeer)
	p.Fail(refusedConnection())
	g := peer.New(
		"foxtrot",
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
			g.Do(
				t.Context(),
				func(context.Context) error { return refusedConnection() },
			),
		),
	)
	deadline := time.Now().Add(time.Second)

	for time.Now().Before(deadline) && p.Probes() < 2 {
		time.Sleep(time.Millisecond)
	}

	assert.True(t, unavailable.Is(g.Unavailable()))
	p.Fail(nil)

	for time.Now().Before(deadline) && g.Unavailable() != nil {
		time.Sleep(time.Millisecond)
	}

	assert.Nil(t, g.Unavailable())
}

func TestACallerThatStopsLeavesThePeerUp(t *testing.T) {
	g := peer.New(
		"golf",
		"http://server.test:8080",
		new(SwitchablePeer),
		logger.New(context.Background()),
		memory.New(),
	)
	defer g.Close()
	x, cancel := context.WithCancel(t.Context())
	cancel()
	e := g.Do(x, func(y context.Context) error { return y.Err() })
	assert.True(t, errors.Is(e, context.Canceled))
	assert.Nil(t, g.Unavailable())
}
