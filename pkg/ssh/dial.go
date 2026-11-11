package ssh

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/errors/connection"
	systemConstant "github.com/funtimecoding/soil/pkg/system/constant"
	"golang.org/x/crypto/ssh"
	"net"
	"strconv"
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

	address := net.JoinHostPort(c.host, strconv.Itoa(c.port))
	configuration := &ssh.ClientConfig{
		User:            c.user,
		Auth:            []ssh.AuthMethod{method},
		HostKeyCallback: callback,
	}

	if c.secure {
		configuration.HostKeyAlgorithms = hostKeyAlgorithms(callback, address)
	}

	result, g := ssh.Dial(systemConstant.Transmission, address, configuration)

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
