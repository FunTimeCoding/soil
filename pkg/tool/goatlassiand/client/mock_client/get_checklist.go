package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goatlassiand/generated/client"
)

func (c *Client) GetChecklist(
	_ context.Context,
	_ string,
) ([]*client.ChecklistItem, error) {
	return nil, c.Failure
}
