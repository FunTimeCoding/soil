package client

import (
	"github.com/funtimecoding/soil/pkg/console/response"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/client"
)

func (c *Client) ListStorageContent(
	name string,
	storage string,
) *response.Response {
	result, e := c.client.ListStorageContentWithResponse(
		c.context,
		name,
		storage,
		&client.ListStorageContentParams{Instance: &c.instance},
	)
	errors.PanicOnError(e)

	return response.New(string(result.Body), result.StatusCode())
}
