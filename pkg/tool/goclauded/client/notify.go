package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
)

func (c *Client) Notify(
	callsign string,
	source string,
	body string,
) error {
	_, e := c.PostNotify(
		context.Background(),
		client.PostNotifyJSONRequestBody{
			Callsign: callsign,
			Source:   source,
			Body:     body,
		},
	)

	return e
}
