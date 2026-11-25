package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/gitlab/constant"
	"github.com/funtimecoding/soil/pkg/gitlab/note"
	"time"
)

func (c *Client) CreateMergeRequestNote(
	_ context.Context,
	project int64,
	identifier int64,
	body string,
) (*note.Note, error) {
	if c.failure != nil {
		return nil, c.failure
	}

	return note.New(
		c.SeedNote(
			project,
			identifier,
			constant.MockUser,
			body,
			false,
			time.Now(),
		),
	), nil
}
