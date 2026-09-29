package gitlab

import "github.com/funtimecoding/soil/pkg/gitlab/runner"

func (c *Client) enrichRunners(v []*runner.Runner) ([]*runner.Runner, error) {
	for _, r := range v {
		if _, e := c.enrichRunner(r); e != nil {
			return nil, e
		}
	}

	return v, nil
}
