package connector

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
	"net/http"
)

func (c *Client) pulseWith(
	sessionIdentifier string,
	body string,
	d client.Delivery,
) (bool, error) {
	response, e := c.generated.PostSessionPulseWithResponse(
		context.Background(),
		sessionIdentifier,
		client.PostSessionPulseJSONRequestBody{Body: body, Delivery: &d},
	)

	if e != nil {
		return false, e
	}

	if response.StatusCode() != http.StatusOK {
		return false, unexpected.Format(
			"pulse status %d",
			response.StatusCode(),
		)
	}

	return response.JSON200 != nil && response.JSON200.Immediate, nil
}
