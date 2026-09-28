package mock_inventory_source

import "github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"

func (c *Client) AddMachine(
	identifier int32,
	name string,
	labels *[]*client.Label,
) {
	c.machines = append(
		c.machines,
		&client.VirtualMachine{
			Identifier: identifier,
			Name:       name,
			Labels:     labels,
		},
	)
}
