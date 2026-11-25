package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	gitlabConstant "github.com/funtimecoding/soil/pkg/gitlab/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/constant"
	"github.com/funtimecoding/soil/pkg/tool/gogitlabd/unit/base"
	"strings"
	"testing"
	"time"
)

func TestListMergeRequestsUpdatedAfterDuration(t *testing.T) {
	s := base.New(t)
	s.Alfa().SeedProject(1, "group", "app")
	s.Alfa().SeedRequest(
		1,
		1,
		gitlabConstant.OpenedState,
		"recent change",
		"delta",
		time.Now().Add(-48*time.Hour),
	)
	s.Alfa().SeedRequest(
		1,
		2,
		gitlabConstant.OpenedState,
		"old change",
		"delta",
		time.Now().Add(-240*time.Hour),
	)
	result := s.Tools().MustCallTool(
		constant.ListMergeRequests,
		map[string]any{"project": "group/app", "updated_after": "7d"},
	)
	assert.StringContains(t, "recent change", result)
	assert.False(t, strings.Contains(result, "old change"))
}

func TestListMergeRequestsUnreadableUpdatedAfterQuiet(t *testing.T) {
	s := base.New(t)
	s.Alfa().SeedProject(1, "group", "app")
	assert.StringContains(
		t,
		"updated_after is neither a duration (7d, 36h) nor RFC3339: yesterday",
		s.Tools().MustCallToolError(
			constant.ListMergeRequests,
			map[string]any{"project": "group/app", "updated_after": "yesterday"},
		),
	)
	assert.Count(t, 0, s.Reporter.Events())
}

func TestListMergeRequestsDefaultLimit(t *testing.T) {
	s := base.New(t)
	s.Alfa().SeedProject(1, "group", "app")
	now := time.Now()

	for i := range int64(12) {
		s.Alfa().SeedRequest(
			1,
			i+1,
			gitlabConstant.OpenedState,
			"change",
			"delta",
			now,
		)
	}

	result := s.Tools().MustCallTool(
		constant.ListMergeRequests,
		map[string]any{"project": "group/app"},
	)
	assert.Integer(t, 10, strings.Count(result, `"Identifier"`))
}
