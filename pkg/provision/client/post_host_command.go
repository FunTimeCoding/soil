package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/connection"
	"github.com/funtimecoding/soil/pkg/provision/generated/client"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) PostHostCommand(
	x context.Context,
	command string,
	host string,
	source string,
) (client.HostCommandRun, error) {
	response, e := c.client.PostHostCommandWithResponse(
		x,
		client.HostCommandRequest{Command: command, Host: host, Source: source},
	)

	if e != nil {
		if f := connection.Classify(e); f != nil {
			return client.HostCommandRun{}, f
		}

		return client.HostCommandRun{}, e
	}

	if response.JSON200 == nil {
		return client.HostCommandRun{}, refusal.New(
			"post host command",
			response.HTTPResponse,
			response.Body,
		)
	}

	return *response.JSON200, nil
}
