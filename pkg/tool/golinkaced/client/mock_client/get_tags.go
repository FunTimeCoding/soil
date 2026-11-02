package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/generated/client"
)

func (c *Client) GetTags(
	_ context.Context,
	_ *client.GetTagsParams,
) (*client.TagPage, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.Failure != nil {
		return nil, c.Failure
	}

	return new(client.TagPage), nil
}
