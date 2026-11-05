package chromium

import (
	"net"
	"strconv"
)

func (c *Client) Address() string {
	return net.JoinHostPort(c.host, strconv.Itoa(c.port))
}
