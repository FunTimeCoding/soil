package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/gobandd/generated/client"
	"maps"
	"slices"
)

func (c *Client) ListInstances(_ context.Context) ([]client.Instance, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	if e := c.failures.Caller(); e != nil {
		return nil, e
	}

	var result []client.Instance

	for _, name := range slices.Sorted(maps.Keys(c.statuses)) {
		result = append(
			result,
			client.Instance{Name: name, Host: c.statuses[name].Host},
		)
	}

	return result, nil
}
