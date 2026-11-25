package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/gitlab/merge_request_filter"
	"github.com/funtimecoding/soil/pkg/gitlab/unit/request_server"
	"testing"
	"time"
)

func TestMergeRequestsQuery(t *testing.T) {
	s := request_server.New(t, 1)
	f := merge_request_filter.New(7)
	since := time.Date(2026, 11, 18, 0, 0, 0, 0, time.UTC)
	draft := true
	f.UpdatedAfter = &since
	f.Author = "alfa"
	f.Reviewer = "bravo"
	f.Assignee = "charlie"
	f.SourceBranch = "feature"
	f.TargetBranch = "main"
	f.Search = "change"
	f.Draft = &draft
	result := s.List(f)
	q := s.Queries()
	assert.Count(t, 1, q)
	assert.String(t, "opened", q[0].Get("state"))
	assert.String(t, "updated_at", q[0].Get("order_by"))
	assert.String(t, "desc", q[0].Get("sort"))
	assert.String(t, "10", q[0].Get("per_page"))
	assert.String(t, "1", q[0].Get("page"))
	assert.String(t, "2026-11-18T00:00:00Z", q[0].Get("updated_after"))
	assert.String(t, "alfa", q[0].Get("author_username"))
	assert.String(t, "bravo", q[0].Get("reviewer_username"))
	assert.String(t, "charlie", q[0].Get("assignee_username[]"))
	assert.String(t, "feature", q[0].Get("source_branch"))
	assert.String(t, "main", q[0].Get("target_branch"))
	assert.String(t, "change", q[0].Get("search"))
	assert.String(t, "true", q[0].Get("draft"))
	assert.Count(t, 1, result)
	r := result[0]
	assert.String(t, "alfa", r.Author)
	assert.Any(t, []string{"bravo"}, r.Reviewers)
	assert.Any(t, []string{"charlie"}, r.Assignees)
	assert.String(t, "feature", r.SourceBranch)
	assert.String(t, "main", r.TargetBranch)
	assert.True(t, r.Draft)
	assert.String(t, "ci_still_running", r.MergeStatus)
	assert.Any(t, int64(3), r.Notes)
	assert.String(t, "2026-11-25T10:00:00Z", r.Update.Format(time.RFC3339))
}

func TestMergeRequestsOffsetAcrossPages(t *testing.T) {
	s := request_server.New(t, 10)
	f := merge_request_filter.New(7)
	f.Limit = 3
	f.Offset = 4
	assert.Any(t, []int64{5, 6, 7}, request_server.Identifiers(s.List(f)))
	q := s.Queries()
	assert.Count(t, 2, q)
	assert.String(t, "2", q[0].Get("page"))
	assert.String(t, "3", q[1].Get("page"))
	assert.String(t, "3", q[1].Get("per_page"))
}

func TestMergeRequestsAlignedOffsetOnePage(t *testing.T) {
	s := request_server.New(t, 10)
	f := merge_request_filter.New(7)
	f.Limit = 3
	f.Offset = 3
	assert.Any(t, []int64{4, 5, 6}, request_server.Identifiers(s.List(f)))
	assert.Count(t, 1, s.Queries())
}

func TestMergeRequestsShortPageEnds(t *testing.T) {
	s := request_server.New(t, 5)
	f := merge_request_filter.New(7)
	f.Limit = 3
	f.Offset = 4
	assert.Any(t, []int64{5}, request_server.Identifiers(s.List(f)))
	assert.Count(t, 1, s.Queries())
}

func TestMergeRequestsOffsetBeyondEnd(t *testing.T) {
	f := merge_request_filter.New(7)
	f.Limit = 3
	f.Offset = 4
	assert.Count(t, 0, request_server.New(t, 2).List(f))
}
