package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/gitlab/unit/note_server"
	"testing"
)

func TestMergeRequestNotesAfterOldestFirst(t *testing.T) {
	s := note_server.New(t, 12)
	assert.Any(t, []int64{10, 11, 12}, note_server.Identifiers(s.Notes(9)))
	q := s.Queries()
	assert.Count(t, 1, q)
	assert.String(t, "created_at", q[0].Get("order_by"))
	assert.String(t, "desc", q[0].Get("sort"))
	assert.String(t, "100", q[0].Get("per_page"))
	assert.String(t, "1", q[0].Get("page"))
}

func TestMergeRequestNotesPagesUntilSeen(t *testing.T) {
	s := note_server.New(t, 250)
	result := s.Notes(120)
	assert.Count(t, 130, result)
	assert.Any(t, int64(121), result[0].Identifier)
	assert.Any(t, int64(250), result[len(result)-1].Identifier)
	assert.Count(t, 2, s.Queries())
}

func TestMergeRequestNotesShortPageEnds(t *testing.T) {
	s := note_server.New(t, 30)
	assert.Count(t, 30, s.Notes(0))
	assert.Count(t, 1, s.Queries())
}

func TestMergeRequestNotesPageLimit(t *testing.T) {
	s := note_server.New(t, 800)
	result := s.Notes(0)
	assert.Count(t, 500, result)
	assert.Any(t, int64(301), result[0].Identifier)
	assert.Count(t, 5, s.Queries())
}

func TestMergeRequestApprovalsSorted(t *testing.T) {
	assert.Strings(
		t,
		[]string{"delta", "echo"},
		note_server.New(t, 0).Approvals(),
	)
}
