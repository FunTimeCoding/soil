package client

import (
	"github.com/funtimecoding/soil/pkg/console/response"
	"github.com/funtimecoding/soil/pkg/errors"
)

func (c *Client) ListInstances() *response.Response {
	result, e := c.client.ListInstancesWithResponse(c.context)
	errors.PanicOnError(e)

	return response.New(string(result.Body), result.StatusCode())
}
