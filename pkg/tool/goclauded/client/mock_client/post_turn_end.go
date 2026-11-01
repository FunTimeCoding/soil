package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
)

func (c *Client) PostTurnEnd(_ context.Context, _ client.PostTurnEndJSONRequestBody) error {
	return c.Failure
}
