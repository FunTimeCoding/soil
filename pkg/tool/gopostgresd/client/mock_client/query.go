package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/generated/client"
)

func (c *Client) Query(
	_ context.Context,
	_ client.QueryJSONRequestBody,
) (*client.QueryResult, error) {
	return c.answer()
}
