package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
)

func (c *Client) PostAnnounce(_ context.Context, _ client.PostAnnounceJSONRequestBody) error {
	return c.Failure
}
