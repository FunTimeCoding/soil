package mock_client

import "github.com/funtimecoding/soil/pkg/tool/goclauded/types/target"

func (c *Client) RecentSessions(limit int) ([]*target.Target, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.Failure != nil {
		return nil, c.Failure
	}

	return c.sessions[:min(limit, len(c.sessions))], nil
}
