package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
)

func (c *Client) PostSessionContext(_ context.Context, _ string, _ client.PostSessionContextJSONRequestBody) error {
	return c.Failure
}
