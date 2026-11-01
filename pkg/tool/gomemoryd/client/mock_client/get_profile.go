package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/generated/client"
)

func (c *Client) GetProfile(
	context.Context,
	*client.GetProfileParams,
) (*client.ProfileResponse, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	return &client.ProfileResponse{}, nil
}
