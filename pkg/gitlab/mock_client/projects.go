package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/gitlab/project"
)

func (c *Client) Projects(x context.Context) ([]*project.Project, error) {
	if e := x.Err(); e != nil {
		return nil, e
	}

	if c.failure != nil {
		return nil, c.failure
	}

	if c.projectFailure != nil {
		return nil, c.projectFailure
	}

	return c.projects, nil
}
