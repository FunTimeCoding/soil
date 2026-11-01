package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
)

func (c *Client) PostRelease(_ context.Context, _ client.PostReleaseJSONRequestBody) error {
	return c.Failure
}
