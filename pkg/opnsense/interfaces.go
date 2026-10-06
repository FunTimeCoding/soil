package opnsense

import (
	"github.com/funtimecoding/soil/pkg/opnsense/constant"
	"github.com/funtimecoding/soil/pkg/opnsense/network_interface"
	"github.com/funtimecoding/soil/pkg/opnsense/response"
	"sort"
)

func (c *Client) Interfaces() ([]*network_interface.Interface, error) {
	var out map[string]response.NetworkInterface

	if e := c.basic.Get(constant.InterfaceState, nil, &out); e != nil {
		return nil, e
	}

	var devices []string

	for device := range out {
		devices = append(devices, device)
	}

	sort.Strings(devices)
	var result []*network_interface.Interface

	for _, device := range devices {
		result = append(result, network_interface.New(device, out[device]))
	}

	return result, nil
}
