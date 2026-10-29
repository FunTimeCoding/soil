package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/client"
)

func (c *Client) DeleteChecklistItem(
	_ context.Context,
	_ string,
	_ int,
) ([]*client.ChecklistItem, error) {
	return nil, c.Failure
}
