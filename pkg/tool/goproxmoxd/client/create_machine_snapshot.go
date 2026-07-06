package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) CreateMachineSnapshot(
	identifier int64,
	name string,
	node *string,
) string {
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

	return web.ReadString(result.HTTPResponse)
}
