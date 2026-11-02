package mock_client

import "github.com/funtimecoding/soil/pkg/tool/gobandd/generated/client"

func (c *Client) SetStatus(s *client.Status) {
	c.statuses[s.Instance] = s
}
