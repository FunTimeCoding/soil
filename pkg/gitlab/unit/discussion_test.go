package unit

import (
	"github.com/funtimecoding/soil/pkg/assert"
	"github.com/funtimecoding/soil/pkg/errors/validation"
	"github.com/funtimecoding/soil/pkg/gitlab/constant"
	"github.com/funtimecoding/soil/pkg/gitlab/unit/discussion_server"
	"testing"
)

func TestReplyToDiscussionPostsIntoThread(t *testing.T) {
	s := discussion_server.New(t)
	n, e := s.Client().ReplyToDiscussion(
		t.Context(),
		7,
		3,
		constant.FixtureDiscussion,
		"see the runbook",
	)
	assert.FatalOnError(t, e)
	assert.Any(t, int64(11), n.Identifier)
	assert.Strings(
		t,
		[]string{
			`POST /api/v4/projects/7/merge_requests/3/discussions/0123456789abcdef0123456789abcdef01234567/notes {"body":"see the runbook"}`,
		},
		s.Requests(),
	)
}

func TestResolveDiscussionPutsResolved(t *testing.T) {
	s := discussion_server.New(t)
	assert.FatalOnError(
		t,
		s.Client().ResolveDiscussion(
			t.Context(),
			7,
			3,
			constant.FixtureDiscussion,
		),
	)
	assert.Strings(
		t,
		[]string{
			`PUT /api/v4/projects/7/merge_requests/3/discussions/0123456789abcdef0123456789abcdef01234567 {"resolved":true}`,
		},
		s.Requests(),
	)
}

func TestDiscussionShapeCheckedBeforeRequest(t *testing.T) {
	s := discussion_server.New(t)

	for _, thread := range []string{
		"",
		"0123456789abcdef",
		"0123456789ABCDEF0123456789ABCDEF01234567",
		"../../../notes?x=0123456789abcdef0123456",
		"0123456789abcdef0123456789abcdef012345670",
	} {
		_, e := s.Client().ReplyToDiscussion(t.Context(), 7, 3, thread, "x")
		assert.True(t, validation.Is(e))
		assert.True(
			t,
			validation.Is(
				s.Client().ResolveDiscussion(t.Context(), 7, 3, thread),
			),
		)
	}

	assert.Count(t, 0, s.Requests())
}
