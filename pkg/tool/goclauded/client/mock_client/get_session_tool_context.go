package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
)

func (c *Client) GetSessionToolContext(
	_ context.Context,
	_ string,
	_ *client.GetSessionToolContextParams,
) (*client.ToolContextResponse, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	return &client.ToolContextResponse{Results: c.ToolContext}, nil
}
