package client

import (
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/provision/generated/client"
	"net/http"
	"strings"
)

func (c *Client) Trigger(changes []string) error {
	response, e := c.client.PostTriggerWithResponse(
		c.context,
		client.TriggerRequest{Update: true, Changes: changes},
	)

	if e != nil {
		return e
	}

	if response.StatusCode() != http.StatusAccepted {
		return unexpected.Format(
			"status %d: %s",
			response.StatusCode(),
			strings.TrimSpace(string(response.Body)),
		)
	}

	return nil
}
