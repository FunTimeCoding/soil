package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/generated/client"
)

func (c *Client) GetThread(
	_ context.Context,
	identifier string,
) ([]*client.Post, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	result, okay := c.Threads[identifier]

	if !okay {
		return nil, not_found.New("post", identifier)
	}

	return result, nil
}
