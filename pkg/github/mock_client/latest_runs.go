package mock_client

import "github.com/funtimecoding/soil/pkg/github/run"

func (c *Client) LatestRuns(
	_ string,
	repository string,
) ([]*run.Run, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	return c.runs[repository], nil
}
