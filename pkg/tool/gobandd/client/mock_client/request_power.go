package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/band"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/generated/client"
)

func (c *Client) RequestPower(
	_ context.Context,
	p client.PowerRequest,
) (*client.PowerResult, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	if e := c.failures.Caller(); e != nil {
		return nil, e
	}

	var name string

	if p.Instance != nil {
		name = *p.Instance
	}

	s, found := c.statuses[name]

	if !found {
		return nil, not_found.Format("instance %s", name)
	}

	state, e := requestedState(string(p.Operation), p.Force != nil && *p.Force)

	if e != nil {
		return nil, e
	}

	c.Requested = append(c.Requested, p)

	return &client.PowerResult{
		Instance:  name,
		Power:     s.Power,
		Requested: band.PowerStateName(state),
	}, nil
}
