package github

import "github.com/funtimecoding/soil/pkg/github/run"

func (c *Client) FailedRuns(verbose bool) ([]*run.Run, error) {
	var result []*run.Run
	runs, e := c.Runs(true, verbose)

	if e != nil {
		return nil, e
	}

	for _, r := range runs {
		for _, j := range r.Jobs {
			if j.Fail() {
				result = append(result, r)
			}
		}
	}

	return result, nil
}
