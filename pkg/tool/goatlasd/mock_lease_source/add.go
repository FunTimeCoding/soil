package mock_lease_source

import "github.com/funtimecoding/soil/pkg/tool/gopnsensed/generated/client"

func (c *Client) Add(
	hostname string,
	address string,
	hardwareAddress string,
	reserved bool,
) {
	c.leases = append(
		c.leases,
		client.Lease{
			Hostname:        hostname,
			Address:         address,
			HardwareAddress: hardwareAddress,
			Reserved:        reserved,
		},
	)
}
