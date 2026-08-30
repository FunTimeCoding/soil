package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/client"
)

func (c *Client) CreateContainerSnapshot(
	identifier int64,
	name string,
	node *string,
) (string, int) {
	body := client.CreateContainerSnapshotJSONRequestBody{Name: name}
	result, e := c.client.CreateContainerSnapshotWithResponse(
		c.context,
		identifier,
		&client.CreateContainerSnapshotParams{
			Instance: &c.instance,
			Node:     node,
		},
		body,
	)
	errors.PanicOnError(e)

	return string(result.Body), result.StatusCode()
}
