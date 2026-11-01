package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/conflict"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
)

func (c *Client) swapLabel(
	sessionIdentifier string,
	key string,
	expect string,
	value string,
	from string,
) (bool, error) {
	_, e := c.PostSessionLabel(
		context.Background(),
		sessionIdentifier,
		client.PostSessionLabelJSONRequestBody{
			Key:    key,
			Value:  &value,
			Expect: &expect,
			From:   &from,
		},
	)

	if conflict.Is(e) {
		return false, nil
	}

	return e == nil, e
}
