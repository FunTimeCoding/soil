package convert

import (
	"github.com/funtimecoding/soil/pkg/netbox/device"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/server"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/store/label"
)

func LabeledDevices(
	v []*device.Device,
	grouped map[int32][]*label.Label,
) []*server.Device {
	result := make([]*server.Device, 0, len(v))

	for _, d := range v {
		entry := Device(d)

		if labels, okay := grouped[d.Identifier]; okay {
			entry.Labels = new(Labels(labels))
		}

		result = append(result, entry)
	}

	return result
}
