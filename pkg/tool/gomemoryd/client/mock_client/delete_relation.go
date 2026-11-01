package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/generated/client"
)

func (c *Client) DeleteRelation(
	context.Context,
	*client.DeleteRelationParams,
) error {
	return c.Failure
}
