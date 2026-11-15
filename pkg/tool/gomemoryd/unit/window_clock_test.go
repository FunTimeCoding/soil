package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/unit/base"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/unit/window_tester"
	"testing"
	"time"
)

func TestImpressionWindowRunsOnServiceClock(t *testing.T) {
	s := base.NewWith(t, window_tester.Now)
	_, e := s.Store().CreateImpression(
		"clocked impression",
		"tester",
		window_tester.Now().Add(-time.Hour),
	)
	assert.FatalOnError(t, e)
	assert.StringContains(
		t,
		"clocked impression",
		s.Body(t, constant.ImpressionsPath),
	)
}

func TestDashboardImpressionsRunOnServiceClock(t *testing.T) {
	s := base.NewWith(t, window_tester.Now)
	_, e := s.Store().CreateImpression(
		"clocked impression",
		"tester",
		window_tester.Now().Add(-time.Hour),
	)
	assert.FatalOnError(t, e)
	assert.StringContains(
		t,
		"clocked impression",
		s.Body(t, constant.DashboardPath),
	)
}
