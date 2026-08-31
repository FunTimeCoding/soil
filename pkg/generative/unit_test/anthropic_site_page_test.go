package unit_test

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/assert/fixture"
	"github.com/funtimecoding/soil/pkg/generative/anthropic/site/page"
	"testing"
)

func TestPageParse(t *testing.T) {
	result := page.Parse(fixture.Read("claude", "usage-page.html"))
	assert.NotNil(t, result)
	assert.Integer(t, 32, result.SessionPercent)
	assert.String(t, "in 50 min", result.SessionReset)
	assert.Integer(t, 22, result.WeeklyAllPercent)
	assert.String(t, "Wed 8:59 PM", result.WeeklyAllReset)
	assert.Integer(t, 34, result.FablePercent)
	assert.String(t, "Wed 8:59 PM", result.FableReset)
}

func TestPageParseEmpty(t *testing.T) {
	assert.Nil(t, page.Parse(""))
}
