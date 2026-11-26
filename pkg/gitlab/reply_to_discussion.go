package gitlab

import (
	"context"
	"github.com/funtimecoding/soil/pkg/gitlab/discussion"
	"github.com/funtimecoding/soil/pkg/gitlab/note"
	"gitlab.com/gitlab-org/api/client-go/v3"
)

func (c *Client) ReplyToDiscussion(
	x context.Context,
	project int64,
	identifier int64,
	thread string,
	body string,
) (*note.Note, error) {
	if e := discussion.Check(thread); e != nil {
		return nil, e
	}

	result, _, e := c.client.Discussions.AddMergeRequestDiscussionNote(
		project,
		identifier,
		thread,
		&gitlab.AddMergeRequestDiscussionNoteOptions{Body: &body},
		gitlab.WithContext(x),
	)

	if e != nil {
		return nil, wrapError(e)
	}

	return note.New(result), nil
}
