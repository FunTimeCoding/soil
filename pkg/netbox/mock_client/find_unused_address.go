package mock_client

import "net"

func (c *Client) FindUnusedAddress(_ *net.IPNet) (net.IP, error) {
	return nil, c.Failure
}
