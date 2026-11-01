package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
	"net/http"
)

func (c *Client) GetChannelCallsign(
	x context.Context,
	params *client.GetChannelCallsignParams,
) (*client.ChannelCallsignResponse, bool, error) {
	r, e := c.generated.GetChannelCallsignWithResponse(x, params)

	if e != nil {
		return nil, false, e
	}

	if r.StatusCode() == http.StatusNoContent {
		return nil, false, nil
	}

	if r.JSON200 == nil {
		return nil, false, refusal.New(
			"get channel callsign",
			r.HTTPResponse,
			r.Body,
		)
	}

	return r.JSON200, true, nil
}
