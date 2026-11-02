package gitlab

import (
	"context"
	"github.com/funtimecoding/soil/pkg/gitlab/runner"
)

func (c *Client) enrichRunners(
	x context.Context,
	v []*runner.Runner,
) ([]*runner.Runner, error) {
	for _, r := range v {
		if _, e := c.enrichRunner(x, r); e != nil {
			return nil, e
		}
	}

	return v, nil
}
