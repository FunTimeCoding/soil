package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
)

func (c *Client) ListDeviceJournalEntries(_ context.Context, _ string, _ int32, _ int32) ([]*client.JournalEntry, error) {
	return nil, c.Failure
}
