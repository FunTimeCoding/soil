package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/tool/godashboardd/unit/service_tester"
	"testing"
)

func TestUnavailablePeerIsHeldOffWithoutBeingAsked(t *testing.T) {
	s := service_tester.New(t, service_tester.WithUsage())
	s.Usage.Failure = unavailable.Format("nextcloud restarting")
	v := s.Service()
	first := v.Refresh()
	assert.Count(t, 1, first)
	assert.True(t, unavailable.Is(first[0]))
	s.Usage.Failure = nil
	second := v.Refresh()
	assert.Count(t, 1, second)
	assert.True(t, unavailable.Is(second[0]))
	assert.Count(t, 0, s.Reporter.Events())
}
