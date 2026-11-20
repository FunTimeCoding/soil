package mock_client

import "github.com/funtimecoding/soil/pkg/tool/gosecurityd/generated/client"

func (c *Client) findFlag(name string) *client.Flag {
	for _, f := range c.flags {
		if f.Identifier == name {
			return f
		}
	}

	return nil
}
