package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
)

func (c *Client) PostCompaction(_ context.Context, _ client.PostCompactionJSONRequestBody) error {
	return c.Failure
}
