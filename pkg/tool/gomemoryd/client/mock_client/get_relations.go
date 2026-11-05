package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/generated/client"
)

func (c *Client) GetRelations(
	_ context.Context,
	p *client.GetRelationsParams,
) ([]client.Relation, error) {
	c.RelationCalls = append(c.RelationCalls, p)

	if c.Failure != nil {
		return nil, c.Failure
	}

	return c.Edges, nil
}
