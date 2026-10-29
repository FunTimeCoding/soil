package mock_client

import (
	"github.com/netbox-community/go-netbox/v4"
	"net"
)

func (c *Client) CreateInternet(
	_ string,
	_ int64,
	_ net.IP,
	_ net.IPMask,
) (*netbox.IPAddress, error) {
	return nil, c.Failure
}
