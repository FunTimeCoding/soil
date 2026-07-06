package gitlab

import "github.com/funtimecoding/soil/pkg/system"

func (c *Client) Trace(
	project int64,
	job int64,
) (string, error) {
	result, _, e := c.client.Jobs.GetTraceFile(project, job)

	if e != nil {
		return "", e
	}

	return string(system.ReadAll(result)), nil
}
