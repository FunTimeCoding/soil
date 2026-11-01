package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
)

func (c *Client) PostEditSession(_ context.Context, _ client.PostEditSessionJSONRequestBody) error {
	return c.Failure
}
