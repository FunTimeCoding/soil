package mock_client

import "github.com/funtimecoding/soil/pkg/github/run"

func (c *Client) AddRun(
	repository string,
	r *run.Run,
) {
	c.runs[repository] = append(c.runs[repository], r)
}
