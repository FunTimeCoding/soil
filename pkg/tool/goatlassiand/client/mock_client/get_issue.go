package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/client"
)

func (c *Client) GetIssue(
	_ context.Context,
	key string,
	_ bool,
) (*client.JiraIssue, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	for _, i := range c.Found {
		if i.Key == key {
			return i, nil
		}
	}

	return nil, not_found.New("issue", key)
}
