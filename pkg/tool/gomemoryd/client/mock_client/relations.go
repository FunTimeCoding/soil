package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/generated/client"
)

func (c *Client) Relations(context.Context) ([]client.Relation, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	return c.Edges, nil
}
