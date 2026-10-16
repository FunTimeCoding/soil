package client

import (
	"github.com/funtimecoding/soil/pkg/errors/unexpected"
	"github.com/funtimecoding/soil/pkg/provision/generated/client"
	"github.com/funtimecoding/soil/pkg/provision/types/change"
	"net/http"
	"strings"
)

func (c *Client) Trigger(v *change.Change) error {
	hosts := []client.HostEvent{}

	for _, h := range v.Hosts {
		hosts = append(
			hosts,
			client.HostEvent{
				Host:   h.Host,
				Action: client.HostEventAction(h.Action),
			},
		)
	}

	response, e := c.client.PostTriggerWithResponse(
		c.context,
		client.TriggerRequest{
			Update:  true,
			Changes: v.Addresses,
			Hosts:   &hosts,
		},
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
