package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/store/save_option"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/unit/base"
	"testing"
)

func TestDashboardCountsSitInTheSummaryStrip(t *testing.T) {
	s := base.New(t)
	o := save_option.New()
	o.Name = "alfa"
	o.Content = "alfa content"
	o.Type = "feedback"
	o.Tags = []string{"shared"}
	_, e := s.Store().CreateMemory(o)
	assert.FatalOnError(t, e)
	page := body(t, s.Server, constant.DashboardPath, false)
	assert.StringContains(t, "1 active · 0 relations", page)
	assert.StringContains(
		t,
		`<a href="/memories?tag=shared" class="chip">`,
		page,
	)
	assert.StringContains(t, "No recent impressions.", page)
}
