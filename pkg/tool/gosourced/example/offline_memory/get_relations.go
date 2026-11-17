package offline_memory

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/generated/client"
)

func (c *Client) GetRelations(
	context.Context,
	*client.GetRelationsParams,
) ([]client.Relation, error) {
	return nil, nil
}
