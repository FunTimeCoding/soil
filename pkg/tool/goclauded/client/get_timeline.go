package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) GetTimeline(
	x context.Context,
	params *client.GetTimelineParams,
) ([]client.TimelineEntry, error) {
	r, e := c.generated.GetTimelineWithResponse(x, params)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("get timeline", r.HTTPResponse, r.Body)
	}

	return *r.JSON200, nil
}
