package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/client"
)

func (c *Client) EditChecklistItem(
	_ context.Context,
	_ string,
	_ int,
	_ string,
) ([]*client.ChecklistItem, error) {
	return nil, c.Failure
}
