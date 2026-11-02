package gitlab

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/gitlab/project"
)

func (c *Client) FileExists(
	x context.Context,
	p *project.Project,
	branch string,
	file string,
) (bool, error) {
	_, e := c.File(x, p.Identifier, branch, file)

	if e != nil {
		if not_found.Is(e) {
			return false, nil
		}

		return false, e
	}

	return true, nil
}
