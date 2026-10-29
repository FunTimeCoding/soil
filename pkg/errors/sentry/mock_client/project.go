package mock_client

import (
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/errors/sentry/basic/response"
)

func (c *Client) Project(
	_ string,
	projectSlug string,
) (*response.Project, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	for _, p := range c.projects {
		if p.Slug == projectSlug {
			return &p, nil
		}
	}

	return nil, not_found.New("project", projectSlug)
}
