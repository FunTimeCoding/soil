package mock_client

import "github.com/funtimecoding/soil/pkg/tool/gomemoryd/generated/client"

func (c *Client) Relations() []client.Relation {
	return c.Edges
}
