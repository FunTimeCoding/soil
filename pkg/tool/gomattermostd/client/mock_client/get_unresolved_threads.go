package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/generated/client"
)

func (c *Client) GetUnresolvedThreads(
	_ context.Context,
	name string,
	_ *client.GetUnresolvedThreadsParams,
) ([]client.Thread, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	result, okay := c.Unresolved[name]

	if !okay {
		return nil, not_found.New("channel", name)
	}

	return result, nil
}
