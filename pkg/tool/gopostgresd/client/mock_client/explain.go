package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gopostgresd/generated/client"
)

func (c *Client) Explain(
	_ context.Context,
	_ client.ExplainJSONRequestBody,
) (*client.QueryResult, error) {
	return c.answer()
}
