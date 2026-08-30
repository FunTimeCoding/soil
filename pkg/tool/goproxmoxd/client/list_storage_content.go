package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/client"
)

func (c *Client) ListStorageContent(
	name string,
	storage string,
) (string, int) {
	result, e := c.client.ListStorageContentWithResponse(
		c.context,
		name,
		storage,
		&client.ListStorageContentParams{Instance: &c.instance},
	)
	errors.PanicOnError(e)

	return string(result.Body), result.StatusCode()
}
