package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/time/clock_tester"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/model/usage_event"
	"github.com/funtimecoding/soil/pkg/tool/gotelemetryd/unit/base"
	"github.com/funtimecoding/soil/pkg/web"
	"testing"
	"time"
)

func TestHeatmapWindowRunsOnServiceClock(t *testing.T) {
	b := base.New(t, base.WithClock(clock_tester.Window))
	assert.FatalOnError(
		t,
		b.Store().Create(
			&usage_event.Event{
				Tool:      "clocked_tool",
				Surface:   constant.SurfaceWebService,
				Actor:     "tester",
				Outcome:   constant.OutcomeSuccess,
				Kind:      "call",
				CreatedAt: clock_tester.Window().Add(-time.Hour),
			},
		),
	)
	s := b.Served()
	r, e := s.Client().Get(s.Address(constant.HeatmapPath))
	assert.FatalOnError(t, e)
	assert.StringContains(t, "clocked_tool", web.ReadString(r))
}
