package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/generated/client"
)

func (c *Client) ListTables(
	_ context.Context,
	_ *client.ListTablesParams,
) (*client.QueryResult, error) {
	return c.answer()
}
