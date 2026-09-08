package connector

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
	"net/http"
)

func (c *Client) Notify(
	callsign string,
	source string,
	body string,
) error {
	response, e := c.generated.PostNotifyWithResponse(
		context.Background(),
		client.PostNotifyJSONRequestBody{
			Callsign: callsign,
			Source:   source,
			Body:     body,
		},
	)

	if e != nil {
		return e
	}

	if response.StatusCode() != http.StatusOK {
		return unexpected.Format("notify status %d", response.StatusCode())
	}

	return nil
}
