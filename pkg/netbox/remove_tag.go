package netbox

import (
	"github.com/funtimecoding/soil/pkg/console"
	"github.com/funtimecoding/soil/pkg/netbox/device"
)

func (c *Client) RemoveTag(
	deviceName string,
	tag string,
) (*device.Device, error) {
	d, e := c.DeviceByName(deviceName)

	if e != nil {
		return nil, e
	}

	if c.verbose {
		console.Format("remove tag device: %+v\n", d)
		console.Format("remove tag raw device: %+v\n", d.Raw)
	}

	d.RemoveTag(tag)
	w := devicePatch(d)
	tags, f := c.tagsNestedRequest(d.Tags)

	if f != nil {
		return nil, f
	}

	w.SetTags(tags)

	return c.updateDevice(d, w)
}
