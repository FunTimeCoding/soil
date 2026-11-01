package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
)

func (c *Client) GetTimeline(_ context.Context, _ *client.GetTimelineParams) ([]client.TimelineEntry, error) {
	return nil, c.Failure
}
