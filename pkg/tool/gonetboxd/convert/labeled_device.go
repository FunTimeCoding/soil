package convert

import (
	"github.com/funtimecoding/soil/pkg/netbox/device"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/store/label"
)

func LabeledDevice(
	d *device.Device,
	labels []*label.Label,
) *server.Device {
	result := Device(d)
	result.Labels = new(Labels(labels))

	return result
}
