package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/constant"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/generated/client"
	"github.com/funtimecoding/soil/pkg/tool/goalertlogd/unit/base"
	"testing"
	"time"
)

func TestRecordStartRunsOnServiceClock(t *testing.T) {
	s := base.NewWith(t, windowNow)
	limit := 5
	all, e := s.AlertClient(t).GetAlertsAfter(
		t.Context(),
		&client.GetAlertsAfterParams{Identifier: 0, Limit: &limit},
	)
	assert.FatalOnError(t, e)
	assert.Count(t, 2, all)
	start, f := time.Parse(constant.DateFormat, all[0].Start)
	assert.FatalOnError(t, f)
	assert.String(t, "2026-08-16T20:00:00Z", start.UTC().Format(time.RFC3339))
}

func TestResolveEndRunsOnServiceClock(t *testing.T) {
	s := base.NewWith(t, windowNow)
	s.MockClient.Remove("fp1")
	s.Clock.Poll()
	limit := 5
	all, e := s.AlertClient(t).GetAlertsAfter(
		t.Context(),
		&client.GetAlertsAfterParams{Identifier: 0, Limit: &limit},
	)
	assert.FatalOnError(t, e)
	assert.String(t, "fp1", all[0].Fingerprint)
	assert.NotNil(t, all[0].End)
	end, f := time.Parse(constant.DateFormat, *all[0].End)
	assert.FatalOnError(t, f)
	assert.String(t, "2026-08-16T20:00:00Z", end.UTC().Format(time.RFC3339))
}
