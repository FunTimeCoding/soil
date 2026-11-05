package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gomemoryd/generated/client"
)

func (c *Client) GetProfile(
	_ context.Context,
	p *client.GetProfileParams,
) (*client.ProfileResponse, error) {
	c.ProfileCalls = append(c.ProfileCalls, p)

	if c.Failure != nil {
		return nil, c.Failure
	}

	if c.Profile != nil {
		return c.Profile, nil
	}

	return &client.ProfileResponse{}, nil
}
