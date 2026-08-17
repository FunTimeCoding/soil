package opnsense

import (
	"encoding/json"
	"github.com/funtimecoding/soil/pkg/opnsense/constant"
	"github.com/funtimecoding/soil/pkg/opnsense/network_interface"
	"github.com/funtimecoding/soil/pkg/opnsense/response"
	"sort"
)

func (c *Client) Interfaces() ([]*network_interface.Interface, error) {
	b, e := c.basic.Get(constant.InterfaceState, nil)

	if e != nil {
		return nil, e
	}

	var out map[string]response.NetworkInterface

	if f := json.Unmarshal(b, &out); f != nil {
		return nil, f
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
