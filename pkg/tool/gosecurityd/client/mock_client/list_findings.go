package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/generated/client"
)

func (c *Client) ListFindings(
	_ context.Context,
	p *client.ListFindingsParams,
) ([]*client.FindingGroup, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.failure != nil {
		return nil, c.failure
	}

	lowest := 0.0

	if p != nil {
		lowest = convert.Minimum(p.Score)
	}

	result := []*client.FindingGroup{}

	for _, g := range c.groups {
		if convert.Minimum(g.Score) >= lowest {
			result = append(result, g)
		}
	}

	return result, nil
}
