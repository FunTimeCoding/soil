package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
)

func (c *Client) AddVirtualJournalEntry(_ context.Context, _ string, _ string, _ string) (client.JournalEntry, error) {
	return client.JournalEntry{}, c.Failure
}
