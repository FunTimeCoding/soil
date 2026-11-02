package ssh

import (
	"fmt"
	"golang.org/x/crypto/ssh"
)

func (c *Client) dialed() (*ssh.Client, error) {
	if c.client == nil {
		return nil, fmt.Errorf("client not dialed: %s", c.host)
	}

	return c.client, nil
}
