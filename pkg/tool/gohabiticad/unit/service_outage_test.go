package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/tool/gohabiticad/unit/service_tester"
	"github.com/funtimecoding/soil/pkg/web/outage_tester"
	"testing"
)

func TestOutageMarksGuardDownWithoutEarlyCapture(t *testing.T) {
	s := service_tester.New(t)
	s.Source.Failure = outage_tester.RefusedDial()
	v := s.Service()
	_, e := v.UserStats()
	assert.True(t, unavailable.Is(e))
	_, f := v.Tags()
	assert.True(t, unavailable.Is(f))
	assert.Count(t, 0, s.Reporter.Events())
}
