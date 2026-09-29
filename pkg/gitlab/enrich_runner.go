package gitlab

import "github.com/funtimecoding/soil/pkg/gitlab/runner"

func (c *Client) enrichRunner(r *runner.Runner) (*runner.Runner, error) {
	g, e := c.GraphRunner(r.Identifier)

	if e != nil {
		return nil, e
	}

	if n := g.Payload.Runner.Managers.Nodes; len(n) > 0 {
		r.Address = n[0].IPAddress
	}

	r.Validate()

	return r, nil
}
