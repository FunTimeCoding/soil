package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
)

func (c *Client) PostSessionEnd(_ context.Context, _ client.PostSessionEndJSONRequestBody) error {
	return c.Failure
}
