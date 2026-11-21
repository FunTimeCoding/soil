package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/convert"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/generated/client"
	"maps"
	"slices"
)

func (c *Client) ListSystems(
	_ context.Context,
	p *client.ListSystemsParams,
) ([]*client.SystemSummary, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.failure != nil {
		return nil, c.failure
	}

	lowest := 0.0

	if p != nil {
		lowest = convert.Minimum(p.Score)
	}

	names := slices.Collect(maps.Keys(c.snapshots))

	for name := range c.details {
		if !slices.Contains(names, name) {
			names = append(names, name)
		}
	}

	slices.Sort(names)
	result := []*client.SystemSummary{}

	for _, name := range names {
		if s := &c.detail(name).Summary; s.TopScore >= lowest {
			result = append(result, s)
		}
	}

	return result, nil
}
