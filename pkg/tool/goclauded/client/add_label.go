package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
)

func (c *Client) AddLabel(
	sessionIdentifier string,
	key string,
	value string,
	from string,
) error {
	_, e := c.PostSessionLabel(
		context.Background(),
		sessionIdentifier,
		client.PostSessionLabelJSONRequestBody{
			Key:    key,
			Value:  &value,
			Action: new(client.Add),
			From:   &from,
		},
	)

	return e
}
