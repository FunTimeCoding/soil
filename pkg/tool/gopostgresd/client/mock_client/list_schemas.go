package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/generated/client"
)

func (c *Client) ListSchemas(
	_ context.Context,
	_ *client.ListSchemasParams,
) (*client.QueryResult, error) {
	return c.answer()
}
