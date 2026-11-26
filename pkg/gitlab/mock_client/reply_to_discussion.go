package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/gitlab/constant"
	"github.com/funtimecoding/soil/pkg/gitlab/discussion"
	"github.com/funtimecoding/soil/pkg/gitlab/note"
	"time"
)

func (c *Client) ReplyToDiscussion(
	_ context.Context,
	project int64,
	identifier int64,
	thread string,
	body string,
) (*note.Note, error) {
	if e := discussion.Check(thread); e != nil {
		return nil, e
	}

	if c.failure != nil {
		return nil, c.failure
	}

	result := c.SeedNote(
		project,
		identifier,
		constant.MockUser,
		body,
		false,
		time.Now(),
	)
	c.threads[thread] = append(c.threads[thread], result)

	return note.New(result), nil
}
