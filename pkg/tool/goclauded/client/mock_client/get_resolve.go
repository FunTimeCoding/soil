package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/constant"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
)

func (c *Client) GetResolve(
	_ context.Context,
	p *client.GetResolveParams,
) (*client.ResolveResponse, []client.ResolveMatch, error) {
	if c.Failure != nil {
		return nil, nil, c.Failure
	}

	if len(c.Matches) > 0 {
		return nil, c.Matches, nil
	}

	if identifier, okay := c.Resolved[p.Query]; okay {
		return &client.ResolveResponse{Identifier: identifier}, nil, nil
	}

	return nil, nil, not_found.New(constant.SessionKind, p.Query)
}
