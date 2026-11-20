package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gosecurityd/generated/client"
	"maps"
	"slices"
)

func (c *Client) ListSystems(context.Context) ([]*client.SystemSummary, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.failure != nil {
		return nil, c.failure
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
		result = append(result, &c.detail(name).Summary)
	}

	return result, nil
}
