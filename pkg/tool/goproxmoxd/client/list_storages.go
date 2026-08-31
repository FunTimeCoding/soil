package client

import (
	"github.com/funtimecoding/soil/pkg/console/response"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/client"
)

func (c *Client) ListStorages(name string) *response.Response {
	result, e := c.client.ListStoragesWithResponse(
		c.context,
		name,
		&client.ListStoragesParams{Instance: &c.instance},
	)
	errors.PanicOnError(e)

	return response.New(string(result.Body), result.StatusCode())
}
