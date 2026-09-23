package connector

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
	"net/http"
)

func (c *Client) Pulse(
	sessionIdentifier string,
	body string,
) (bool, error) {
	response, e := c.generated.PostSessionPulseWithResponse(
		context.Background(),
		sessionIdentifier,
		client.PostSessionPulseJSONRequestBody{Body: body},
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

	return false, nil
}
