package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/generated/client"
)

func (c *Client) GetRelations(
	context.Context,
	*client.GetRelationsParams,
) ([]client.Relation, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	return c.Edges, nil
}
