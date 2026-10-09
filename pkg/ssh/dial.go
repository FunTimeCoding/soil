package ssh

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors/connection"
	"github.com/funtimecoding/soil/pkg/system/constant"
	"golang.org/x/crypto/ssh"
)

func (c *Client) Dial() error {
	callback, e := c.hostKeyCallback()

	if e != nil {
		return e
	}

	method, f := c.authenticate()

	if f != nil {
		return f
	}

	result, g := ssh.Dial(
		constant.Transmission,
		fmt.Sprintf("%s:22", c.host),
		&ssh.ClientConfig{
			User:            c.user,
			Auth:            []ssh.AuthMethod{method},
			HostKeyCallback: callback,
		},
	)

	if g != nil {
		if h := connection.Classify(g); h != nil {
			return h
		}

		return fmt.Errorf("dial %s: %w", c.host, g)
	}

	c.client = result

	return nil
}
