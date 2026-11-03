package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
)

func (c *Client) GetPostReactions(
	_ context.Context,
	identifier string,
) ([]string, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	if e := c.failures.Caller(); e != nil {
		return nil, e
	}

	result, okay := c.Reactions[identifier]

	if !okay {
		return nil, not_found.New("post", identifier)
	}

	return result, nil
}
