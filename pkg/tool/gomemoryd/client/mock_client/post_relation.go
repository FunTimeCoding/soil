package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/generated/client"
)

func (c *Client) PostRelation(
	context.Context,
	client.PostRelationJSONRequestBody,
) error {
	return c.Failure
}
