package mock_client

import "github.com/funtimecoding/soil/pkg/provision/generated/client"

func (c *Client) HostCommands() []client.HostCommandRequest {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	return append([]client.HostCommandRequest{}, c.hostCommands...)
}
