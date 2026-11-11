package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/backrest/response"
)

func (c *Client) Stanzas(_ context.Context) ([]*response.Stanza, error) {
	c.Calls++

	if c.Failure != nil {
		return nil, c.Failure
	}

	return c.stanzas, nil
}
