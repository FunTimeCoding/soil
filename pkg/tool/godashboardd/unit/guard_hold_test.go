package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/sentry/reporter/memory"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/unit/service_tester"
	"testing"
)

func TestAnUnavailablePeerIsHeldOffWithoutBeingAsked(t *testing.T) {
	p := newPrometheus(t)
	u := newUsage(t)
	u.Failure = unavailable.Format("nextcloud restarting")
	r := memory.New()
	v := service_tester.New(
		t,
		service_tester.NewOption(t, service_tester.FixtureBoard(), p, nil, u),
		r,
	)
	first := v.Refresh()
	assert.Count(t, 1, first)
	assert.True(t, unavailable.Is(first[0]))
	u.Failure = nil
	second := v.Refresh()
	assert.Count(t, 1, second)
	assert.True(t, unavailable.Is(second[0]))
	assert.Count(t, 0, r.Events())
}
