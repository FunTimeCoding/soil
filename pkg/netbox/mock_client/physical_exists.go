package mock_client

import "net"

func (c *Client) PhysicalExists(_ net.HardwareAddr) (bool, error) {
	return false, c.Failure
}
