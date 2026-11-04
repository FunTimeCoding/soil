package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
)

func (c *Client) Notify(
	callsign string,
	source string,
	body string,
	d constant.Delivery,
) error {
	delivery := client.Delivery(d)
	_, e := c.PostNotify(
		context.Background(),
		client.PostNotifyJSONRequestBody{
			Callsign: callsign,
			Source:   source,
			Body:     body,
			Delivery: &delivery,
		},
	)

	return e
}
