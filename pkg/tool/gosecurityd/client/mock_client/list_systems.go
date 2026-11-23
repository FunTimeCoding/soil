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

	if p == nil {
		p = &client.ListSystemsParams{}
	}

	s, e := convert.Selection(p.Show, p.Label, p.Search, p.Score, 0)

	if e != nil {
		return nil, e
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
		if c.chosen(name, s) {
			result = append(result, &c.detail(name).Summary)
		}
	}

	return result, nil
}
