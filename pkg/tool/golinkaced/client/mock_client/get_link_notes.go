package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/generated/client"
)

func (c *Client) GetLinkNotes(
	_ context.Context,
	_ int32,
	_ *client.GetLinkNotesParams,
) (*client.NotePage, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.Failure != nil {
		return nil, c.Failure
	}

	return new(client.NotePage), nil
}
