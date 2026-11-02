package gitlab

import (
	"context"
	"github.com/funtimecoding/soil/pkg/gitlab/runner"
)

func (c *Client) enrichRunner(
	x context.Context,
	r *runner.Runner,
) (*runner.Runner, error) {
	g, e := c.GraphRunner(x, r.Identifier)

	if e != nil {
		return nil, e
	}

	if n := g.Payload.Runner.Managers.Nodes; len(n) > 0 {
		r.Address = n[0].IPAddress
	}

	r.Validate()

	return r, nil
}
