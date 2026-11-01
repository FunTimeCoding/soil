package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/assert/fixture"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/site/page"
	"github.com/funtimecoding/soil/pkg/generative/constant"
	"strings"
	"testing"
)

func TestPageParse(t *testing.T) {
	result := page.Parse(fixture.Read("claude", "usage-page.html"))
	assert.NotNil(t, result)
	assert.True(t, result.SessionSeen)
	assert.Integer(t, 7, result.SessionPercent)
	assert.String(t, "at 4:15 AM", result.SessionReset)
	assert.Integer(t, 41, result.WeeklyAllPercent)
	assert.String(t, "Thursday 6:00 PM", result.WeeklyAllReset)
	assert.True(t, result.FableSeen)
	assert.Integer(t, 18, result.FablePercent)
	assert.String(t, "Thursday 6:00 PM", result.FableReset)
}

func TestPageParseWeekOnly(t *testing.T) {
	result := page.Parse(fixture.Read("claude", "usage-page-week.html"))
	assert.FatalNotNil(t, result)
	assert.Integer(t, 23, result.WeeklyAllPercent)
	assert.String(t, "Wednesday 8:00 PM", result.WeeklyAllReset)
	assert.False(t, result.SessionSeen)
	assert.False(t, result.FableSeen)
}

func TestPageParseWeekOpenReadsFable(t *testing.T) {
	result := page.Parse(fixture.Read("claude", "usage-page-week-open.html"))
	assert.FatalNotNil(t, result)
	assert.Integer(t, 24, result.WeeklyAllPercent)
	assert.String(t, "Wednesday 8:00 PM", result.WeeklyAllReset)
	assert.False(t, result.SessionSeen)
	assert.True(t, result.FableSeen)
	assert.Integer(t, 2, result.FablePercent)
	assert.String(t, "Wednesday 8:00 PM", result.FableReset)
}

func TestPageParseWithoutWeekIsNil(t *testing.T) {
	assert.Nil(
		t,
		page.Parse(
			strings.ReplaceAll(
				fixture.Read("claude", "usage-page.html"),
				constant.UsageMeterWeek,
				"Last week",
			),
		),
	)
}

func TestPageParseEmpty(t *testing.T) {
	assert.Nil(t, page.Parse(""))
}
