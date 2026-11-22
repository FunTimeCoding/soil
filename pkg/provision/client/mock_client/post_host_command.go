package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/provision/generated/client"
)

func (c *Client) PostHostCommand(
	_ context.Context,
	command string,
	host string,
	source string,
) (client.HostCommandRun, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.calls++

	if e := c.failures.Caller(); e != nil {
		return client.HostCommandRun{}, e
	}

	c.hostCommands = append(
		c.hostCommands,
		client.HostCommandRequest{Command: command, Host: host, Source: source},
	)

	return client.HostCommandRun{
		Identifier: len(c.hostCommands),
		Status:     c.hostCommandStatus,
	}, nil
}
