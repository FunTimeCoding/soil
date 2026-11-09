package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/types/usage_window"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/service_tester"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/web/usage_page"
	"strings"
	"testing"
	"time"
)

func TestUsageWindowNowRunsOnServiceClock(t *testing.T) {
	s := service_tester.New(t)
	s.Store.Advance(usageNow().Sub(s.Store.Clock()()))
	assert.FatalOnError(
		t,
		s.Usage.RecordRateLimits(
			0,
			20,
			time.Time{},
			usageNow().Add(24*time.Hour),
		),
	)
	w := s.Usage.UsageWindow()
	assert.NotNil(t, w)
	assert.String(t, "2026-11-09T12:00:00Z", w.Now.UTC().Format(time.RFC3339))
}

func TestUsageChartWithoutNowDrawsNoNowMarker(t *testing.T) {
	w := usage_window.New(
		usageNow().Add(-7*24*time.Hour),
		usageNow(),
		time.Time{},
		nil,
		nil,
	)
	var markup strings.Builder
	assert.FatalOnError(t, usage_page.Chart(w).Render(&markup))
	assert.StringNotContains(t, "chart-now", markup.String())
}
