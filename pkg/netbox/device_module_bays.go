package netbox

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/netbox/module_bay"
	"github.com/funtimecoding/soil/pkg/netbox/refusal"
)

func (c *Client) DeviceModuleBays(device string) ([]*module_bay.Bay, error) {
	console.Format("Name: %s\n", device)
	result, _, e := c.client.DcimAPI.DcimModuleBaysList(c.context).Device(
		[]*string{&device},
	).Execute()

	if e != nil {
		return nil, refusal.Classify(e)
	}

	return module_bay.NewSlice(result.Results), nil
}
