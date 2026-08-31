package client

import (
	"github.com/funtimecoding/soil/pkg/console/response"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/client"
)

func (c *Client) CreateMachineSnapshot(
	identifier int64,
	name string,
	node *string,
) *response.Response {
	body := client.CreateMachineSnapshotJSONRequestBody{Name: name}
	result, e := c.client.CreateMachineSnapshotWithResponse(
		c.context,
		identifier,
		&client.CreateMachineSnapshotParams{
			Instance: &c.instance,
			Node:     node,
		},
		body,
	)
	errors.PanicOnError(e)

	return response.New(string(result.Body), result.StatusCode())
}
