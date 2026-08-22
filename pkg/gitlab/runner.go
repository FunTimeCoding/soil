package gitlab

import "github.com/funtimecoding/soil/pkg/gitlab/runner"

func (c *Client) Runner(identifier int64) (*runner.Runner, error) {
	result, _, e := c.client.Runners.GetRunnerDetails(identifier)

	if e != nil {
		return nil, wrapError(e)
	}

	return c.enrichRunner(runner.NewDetail(result)), nil
}
