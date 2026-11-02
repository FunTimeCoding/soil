package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/generated/client"
)

func (c *Client) DescribeTable(
	_ context.Context,
	_ string,
	_ *client.DescribeTableParams,
) (*client.QueryResult, error) {
	return c.answer()
}
