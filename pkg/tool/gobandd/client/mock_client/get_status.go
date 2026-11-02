package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/generated/client"
)

func (c *Client) GetStatus(
	_ context.Context,
	p *client.GetStatusParams,
) (*client.Status, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	if e := c.failures.Caller(); e != nil {
		return nil, e
	}

	var name string

	if p != nil && p.Instance != nil {
		name = *p.Instance
	}

	result, found := c.statuses[name]

	if !found {
		return nil, not_found.Format("instance %s", name)
	}

	return result, nil
}
