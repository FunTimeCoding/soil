package github

import "github.com/funtimecoding/soil/pkg/github/run"

func (c *Client) FailedRuns(verbose bool) []*run.Run {
	var result []*run.Run

	for _, r := range c.Runs(true, verbose) {
		for _, j := range r.Jobs {
			if j.Fail() {
				result = append(result, r)
			}
		}
	}

	return result
}
