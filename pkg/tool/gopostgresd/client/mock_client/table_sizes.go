package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/generated/client"
)

func (c *Client) TableSizes(
	_ context.Context,
	_ *client.TableSizesParams,
) (*client.QueryResult, error) {
	return c.answer()
}
