package mock_client

import "github.com/funtimecoding/soil/pkg/tool/gosecurityd/generated/client"

func (c *Client) SeedSystemDetail(v *client.SystemDetail) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.details[v.Summary.Name] = v
}
