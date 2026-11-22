package mock_client

import "github.com/funtimecoding/soil/pkg/provision/generated/client"

func (c *Client) SetHostResults(results []client.HostCommandResult) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.hostResults = append([]client.HostCommandResult{}, results...)
}
