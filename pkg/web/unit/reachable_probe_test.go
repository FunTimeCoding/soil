package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"github.com/funtimecoding/soil/pkg/web/reachable_probe"
	"github.com/funtimecoding/soil/pkg/web/unit/guarded_server"
	"testing"
)

func TestProbeReachesGuardedDaemon(t *testing.T) {
	s := guarded_server.New(t)
	assert.FatalOnError(
		t,
		reachable_probe.Reach(
			context.Background(),
			s.Client(),
			locator.NewOrigin(s.URL),
			"alfa",
		),
	)
}

func TestProbeFailsOnWrongToken(t *testing.T) {
	s := guarded_server.New(t)
	e := reachable_probe.Reach(
		context.Background(),
		s.Client(),
		locator.NewOrigin(s.URL),
		"bravo",
	)
	assert.True(t, e != nil)
	assert.StringContains(t, "401", e.Error())
}

func TestProbeFailsWithoutToken(t *testing.T) {
	s := guarded_server.New(t)
	e := reachable_probe.Reach(
		context.Background(),
		s.Client(),
		locator.NewOrigin(s.URL),
		"",
	)
	assert.True(t, e != nil)
}
