package ssh

import (
	"golang.org/x/crypto/ssh"
	"log"
)

func (c *Client) dialed() *ssh.Client {
	if c.client == nil {
		log.Panicf("client not dialed: %s", c.host)
	}

	return c.client
}
