package client

import (
	"github.com/funtimecoding/soil/pkg/console/response"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/client"
)

func (c *Client) DeleteContainerSnapshot(
	identifier int64,
	name string,
	node *string,
) *response.Response {
	result, e := c.client.DeleteContainerSnapshotWithResponse(
		c.context,
		identifier,
		name,
		&client.DeleteContainerSnapshotParams{
			Instance: &c.instance,
			Node:     node,
		},
	)
	errors.PanicOnError(e)

	return response.New(string(result.Body), result.StatusCode())
}
