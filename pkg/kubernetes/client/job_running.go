package client

import "github.com/funtimecoding/soil/pkg/errors/not_found"

func (c *Client) JobRunning(
	namespace string,
	name string,
) (bool, error) {
	j, e := c.Job(namespace, name)

	if not_found.Is(e) {
		return false, nil
	}

	if e != nil {
		return false, e
	}

	return j.Raw.Status.CompletionTime == nil && j.Raw.Status.Failed == 0, nil
}
