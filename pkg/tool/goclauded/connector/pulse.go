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
) error {
	response, e := c.generated.PostSessionPulseWithResponse(
		context.Background(),
		sessionIdentifier,
		client.PostSessionPulseJSONRequestBody{Body: body},
	)

	if e != nil {
		return e
	}

	if response.StatusCode() != http.StatusOK {
		return unexpected.Format("pulse status %d", response.StatusCode())
	}

	return nil
}
