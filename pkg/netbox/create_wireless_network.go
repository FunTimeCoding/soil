package netbox

import (
	"github.com/funtimecoding/soil/pkg/netbox/wireless_network"
	"github.com/netbox-community/go-netbox/v4"
)

func (c *Client) CreateWirelessNetwork(
	ssid string,
) (*wireless_network.Network, error) {
	q := netbox.NewWritableWirelessLANRequest(ssid)
	result, _, e := c.client.WirelessAPI.WirelessWirelessLansCreate(c.context).WritableWirelessLANRequest(
		*q,
	).Execute()

	if e != nil {
		return nil, e
	}

	return wireless_network.New(result), nil
}
