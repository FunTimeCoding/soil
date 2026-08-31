package client

import (
	"github.com/funtimecoding/soil/pkg/console/response"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/client"
)

func (c *Client) CreateContainerSnapshot(
	identifier int64,
	name string,
	node *string,
) *response.Response {
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

	return response.New(string(result.Body), result.StatusCode())
}
