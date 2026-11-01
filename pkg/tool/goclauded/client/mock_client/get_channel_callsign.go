package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
)

func (c *Client) GetChannelCallsign(_ context.Context, _ *client.GetChannelCallsignParams) (*client.ChannelCallsignResponse, bool, error) {
	return &client.ChannelCallsignResponse{}, true, c.Failure
}
