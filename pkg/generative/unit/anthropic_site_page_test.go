package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/assert/fixture"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/site/page"
	"testing"
)

func TestPageParse(t *testing.T) {
	result := page.Parse(fixture.Read("claude", "usage-page.html"))
	assert.NotNil(t, result)
	assert.Integer(t, 7, result.SessionPercent)
	assert.String(t, "at 4:15 AM", result.SessionReset)
	assert.Integer(t, 41, result.WeeklyAllPercent)
	assert.String(t, "Thursday 6:00 PM", result.WeeklyAllReset)
	assert.Integer(t, 18, result.FablePercent)
	assert.String(t, "Thursday 6:00 PM", result.FableReset)
}

func TestPageParseEmpty(t *testing.T) {
	assert.Nil(t, page.Parse(""))
}
