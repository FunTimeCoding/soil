package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gomaintlogd/generated/client"
)

func (c *Client) GetEntries(
	_ context.Context,
	_ *client.GetEntriesParams,
) ([]client.EntryResponse, error) {
	return nil, c.Failure
}
