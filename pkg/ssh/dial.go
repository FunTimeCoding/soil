package ssh

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/connection"
	"github.com/funtimecoding/soil/pkg/system/constant"
	"golang.org/x/crypto/ssh"
)

func (c *Client) Dial() error {
	callback, e := c.hostKeyCallback()

	if e != nil {
		return e
	}

	method, agent, f := c.authenticate()

	if f != nil {
		return f
	}

	if agent != nil {
		defer errors.LogClose(agent)
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
			if h.Host == "" {
				return connection.New(h.Kind, c.host, h.Path, h.Reason)
			}

			return h
		}

		return fmt.Errorf("dial %s: %w", c.host, g)
	}

	c.client = result

	return nil
}
