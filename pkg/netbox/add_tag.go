package netbox

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/netbox/device"
)

func (c *Client) AddTag(
	deviceName string,
	tag string,
) (*device.Device, error) {
	t, e := c.TagByName(tag)

	if e != nil {
		return nil, e
	}

	if c.verbose {
		fmt.Printf(
			"ADD tag: %s %s\n",
			t.Nested.GetName(),
			t.Nested.GetSlug(),
		)
	}

	d, f := c.DeviceByName(deviceName)

	if f != nil {
		return nil, f
	}

	if c.verbose {
		fmt.Printf("add tag device: %+v\n", d)
		fmt.Printf("add tag raw device: %+v\n", d.Raw)
	}

	d.AddTag(t.Name)
	w := devicePatch(d)
	tags, g := c.tagsNestedRequest(d.Tags)

	if g != nil {
		return nil, g
	}

	w.SetTags(tags)

	return c.updateDevice(d, w)
}
