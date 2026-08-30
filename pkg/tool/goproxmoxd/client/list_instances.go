package client

import "github.com/funtimecoding/soil/pkg/errors"

func (c *Client) ListInstances() (string, int) {
	result, e := c.client.ListInstancesWithResponse(c.context)
	errors.PanicOnError(e)

	return string(result.Body), result.StatusCode()
}
