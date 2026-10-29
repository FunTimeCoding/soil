package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
)

func (c *Client) UpdateJournalEntry(_ context.Context, _ int32, _ string, _ string) (client.JournalEntry, error) {
	return client.JournalEntry{}, c.Failure
}
