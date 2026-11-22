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
	request := client.HostCommandRequest{Command: command, Source: source}

	if host != "" {
		request.Host = &host
	}

	response, e := c.client.PostHostCommandWithResponse(x, request)

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
