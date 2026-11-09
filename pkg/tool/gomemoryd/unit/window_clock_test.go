package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/unit/base"
	"testing"
	"time"
)

func TestImpressionWindowRunsOnServiceClock(t *testing.T) {
	s := base.NewWith(t, windowNow)
	_, e := s.Store().CreateImpression(
		"clocked impression",
		"tester",
		windowNow().Add(-time.Hour),
	)
	assert.FatalOnError(t, e)
	assert.StringContains(
		t,
		"clocked impression",
		body(t, s.Server, constant.ImpressionsPath, false),
	)
}

func TestDashboardImpressionsRunOnServiceClock(t *testing.T) {
	s := base.NewWith(t, windowNow)
	_, e := s.Store().CreateImpression(
		"clocked impression",
		"tester",
		windowNow().Add(-time.Hour),
	)
	assert.FatalOnError(t, e)
	assert.StringContains(
		t,
		"clocked impression",
		body(t, s.Server, constant.DashboardPath, false),
	)
}
