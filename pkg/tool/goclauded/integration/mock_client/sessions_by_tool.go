package mock_client

import "github.com/funtimecoding/soil/pkg/generative/anthropic/claude"

func (c *Client) SessionsByTool(toolFilter string) []*claude.SessionToolCount {
	return nil
}
