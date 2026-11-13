package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/log/logger"
	"github.com/funtimecoding/soil/pkg/peer"
	"testing"
)

func TestGuardedHandsBackCallValue(t *testing.T) {
	g := peer.New(
		"echo",
		"http://server.test:8080",
		new(SwitchablePeer),
		logger.New(context.Background()),
		memory.New(),
	)
	defer g.Close()
	result, e := peer.Guarded(
		t.Context(),
		g,
		func(context.Context) (string, error) { return "foxtrot", nil },
	)
	assert.FatalOnError(t, e)
	assert.String(t, "foxtrot", result)
}

func TestGuardedOutageAnswersWithoutReachingPeer(t *testing.T) {
	g := peer.New(
		"golf",
		"http://server.test:8080",
		new(SwitchablePeer),
		logger.New(context.Background()),
		memory.New(),
	)
	defer g.Close()
	calls := 0
	call := func(context.Context) (string, error) {
		calls++

		return "hotel", refusedConnection()
	}
	_, e := peer.Guarded(t.Context(), g, call)
	assert.True(t, unavailable.Is(e))
	result, f := peer.Guarded(t.Context(), g, call)
	assert.True(t, unavailable.Is(f))
	assert.String(t, "", result)
	assert.Integer(t, 1, calls)
}
