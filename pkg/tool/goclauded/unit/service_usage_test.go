package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/assert/fixture"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/site/page"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/unit/service_tester"
	"testing"
	"time"
)

func TestNoUsageYetIsAbsentNotFailure(t *testing.T) {
	s := service_tester.New(t)
	result, found, e := s.Usage.FindUsage()
	assert.FatalOnError(t, e)
	assert.False(t, found)
	assert.True(t, result == nil)
}

func TestFailedUsageReadIsFailureNotAbsence(t *testing.T) {
	s := service_tester.New(t)
	s.Close()
	_, found, e := s.Usage.FindUsage()
	assert.NotNil(t, e)
	assert.False(t, found)
}

func TestUsageHidesFableFromEarlierWeek(t *testing.T) {
	s := service_tester.New(t)
	assert.FatalOnError(t, s.Usage.RecordFable(34, "Wed 8:59 PM", nil))
	s.Store.Advance(8 * 24 * time.Hour)
	now := s.Store.Clock()()
	assert.FatalOnError(
		t,
		s.Usage.RecordRateLimits(0, 20, time.Time{}, now.Add(24*time.Hour)),
	)
	result, found, e := s.Usage.FindUsage()
	assert.FatalOnError(t, e)
	assert.True(t, found)
	assert.False(t, result.HasFable())
	assert.Integer(t, 20, result.SevenDayPercent)
}

func TestUsageShowsFableFromThisWeek(t *testing.T) {
	s := service_tester.New(t)
	now := s.Store.Clock()()
	assert.FatalOnError(
		t,
		s.Usage.RecordRateLimits(0, 20, time.Time{}, now.Add(24*time.Hour)),
	)
	assert.FatalOnError(t, s.Usage.RecordFable(34, "Wed 8:59 PM", nil))
	result, found, e := s.Usage.FindUsage()
	assert.FatalOnError(t, e)
	assert.True(t, found)
	assert.True(t, result.HasFable())
	assert.Integer(t, 34, result.FablePercent)
}

func TestUsageWithoutFiveHourHasNoSession(t *testing.T) {
	s := service_tester.New(t)
	now := s.Store.Clock()()
	assert.FatalOnError(
		t,
		s.Usage.RecordRateLimits(0, 23, time.Time{}, now.Add(24*time.Hour)),
	)
	result, found, e := s.Usage.FindUsage()
	assert.FatalOnError(t, e)
	assert.True(t, found)
	assert.False(t, result.HasSession())
	assert.Integer(t, 23, result.SevenDayPercent)
}

func TestUsagePageWeekOnlyRecordsNoFable(t *testing.T) {
	s := service_tester.New(t)
	assert.FatalOnError(
		t,
		s.Usage.RecordUsagePage(
			page.Parse(fixture.Read("claude", "usage-page-week.html")),
		),
	)
	fable, e := s.Store.Store.LatestFableSnapshot()
	assert.FatalOnError(t, e)
	assert.Nil(t, fable)
	assert.Count(t, 0, s.Reporter.Events())
}

func TestUsagePageRecordsFableWhenSeen(t *testing.T) {
	s := service_tester.New(t)
	assert.FatalOnError(
		t,
		s.Usage.RecordUsagePage(
			page.Parse(fixture.Read("claude", "usage-page.html")),
		),
	)
	fable, e := s.Store.Store.LatestFableSnapshot()
	assert.FatalOnError(t, e)
	assert.FatalNotNil(t, fable)
	assert.Integer(t, 18, fable.Percent)
	assert.String(t, "Thursday 6:00 PM", fable.Reset)
}

func TestUsagePageRecordsFableFromOpenWeek(t *testing.T) {
	s := service_tester.New(t)
	assert.FatalOnError(
		t,
		s.Usage.RecordUsagePage(
			page.Parse(fixture.Read("claude", "usage-page-week-open.html")),
		),
	)
	fable, e := s.Store.Store.LatestFableSnapshot()
	assert.FatalOnError(t, e)
	assert.FatalNotNil(t, fable)
	assert.Integer(t, 2, fable.Percent)
	assert.String(t, "Wednesday 8:00 PM", fable.Reset)
}

func TestUnreadableUsagePageCapturesOncePerEpisode(t *testing.T) {
	s := service_tester.New(t)
	assert.FatalOnError(t, s.Usage.RecordUsagePage(nil))
	assert.FatalOnError(t, s.Usage.RecordUsagePage(nil))
	assert.FatalOnError(t, s.Usage.RecordUsagePage(nil))
	assert.Count(t, 1, s.Reporter.Events())
	assert.FatalOnError(
		t,
		s.Usage.RecordUsagePage(
			page.Parse(fixture.Read("claude", "usage-page-week.html")),
		),
	)
	assert.FatalOnError(t, s.Usage.RecordUsagePage(nil))
	assert.Count(t, 2, s.Reporter.Events())
}
