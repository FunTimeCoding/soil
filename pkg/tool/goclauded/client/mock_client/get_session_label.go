package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
	"maps"
	"slices"
)

func (c *Client) GetSessionLabel(
	_ context.Context,
	identifier string,
) (*client.SessionLabelResponse, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.Failure != nil {
		return nil, c.Failure
	}

	result := &client.SessionLabelResponse{Labels: []client.LabelEntry{}}
	s := c.find(identifier)

	if s == nil {
		return result, nil
	}

	for _, key := range slices.Sorted(maps.Keys(s.Labels)) {
		for _, value := range s.Labels[key] {
			result.Labels = append(
				result.Labels,
				client.LabelEntry{Key: key, Value: value},
			)
		}
	}

	return result, nil
}
