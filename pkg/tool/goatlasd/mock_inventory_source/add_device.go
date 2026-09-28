package mock_inventory_source

import "github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"

func (c *Client) AddDevice(
	identifier int32,
	name string,
	labels *[]*client.Label,
) {
	c.devices = append(
		c.devices,
		&client.Device{Identifier: identifier, Name: name, Labels: labels},
	)
}
