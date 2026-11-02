package client

import (
	"github.com/funtimecoding/soil/pkg/errors/connection"
	"github.com/funtimecoding/soil/pkg/provision/generated/client"
	"github.com/funtimecoding/soil/pkg/provision/types/change"
	"github.com/funtimecoding/soil/pkg/web/refusal"
	"net/http"
)

func (c *Client) PostTrigger(v *change.Change) error {
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
		if f := connection.Classify(e); f != nil {
			return f
		}

		return e
	}

	if response.StatusCode() != http.StatusAccepted {
		return refusal.New("post trigger", response.HTTPResponse, response.Body)
	}

	return nil
}
