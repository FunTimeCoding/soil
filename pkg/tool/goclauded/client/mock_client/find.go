package mock_client

import "github.com/funtimecoding/soil/pkg/tool/goclauded/types/target"

func (c *Client) find(identifier string) *target.Target {
	for _, s := range c.sessions {
		if s.Identifier == identifier {
			return s
		}
	}

	return nil
}
