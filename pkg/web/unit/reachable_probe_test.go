package unit

import (
	"context"
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/web/reachable_probe"
	"github.com/funtimecoding/soil/pkg/web/upstream_tester"
	"testing"
)

func TestTheProbeReachesAGuardedDaemon(t *testing.T) {
	s := reachableServer(t)
	assert.FatalOnError(
		t,
		reachable_probe.Reach(
			context.Background(),
			s.Client(),
			upstream_tester.Locator(t, s.URL),
			"alfa",
		),
	)
}

func TestTheProbeFailsOnAWrongToken(t *testing.T) {
	s := reachableServer(t)
	e := reachable_probe.Reach(
		context.Background(),
		s.Client(),
		upstream_tester.Locator(t, s.URL),
		"bravo",
	)
	assert.True(t, e != nil)
	assert.StringContains(t, "401", e.Error())
}

func TestTheProbeFailsWithoutAToken(t *testing.T) {
	s := reachableServer(t)
	e := reachable_probe.Reach(
		context.Background(),
		s.Client(),
		upstream_tester.Locator(t, s.URL),
		"",
	)
	assert.True(t, e != nil)
}
