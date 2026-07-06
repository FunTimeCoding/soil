package netbox

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/netbox/module_bay"
)

func (c *Client) DeviceModuleBays(device string) ([]*module_bay.Bay, error) {
	fmt.Printf("Name: %s\n", device)
	result, _, e := c.client.DcimAPI.DcimModuleBaysList(
		c.context,
	).Device([]*string{&device}).Execute()

	if e != nil {
		return nil, e
	}

	return module_bay.NewSlice(result.Results), nil
}
