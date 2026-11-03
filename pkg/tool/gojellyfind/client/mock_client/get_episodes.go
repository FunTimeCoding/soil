package mock_client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
	"github.com/funtimecoding/soil/pkg/tool/gojellyfind/generated/client"
	"slices"
)

func (c *Client) GetEpisodes(
	_ context.Context,
	identifier string,
) ([]*client.Item, error) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.Failure != nil {
		return nil, c.Failure
	}

	if e := c.failures.Caller(); e != nil {
		return nil, e
	}

	if identifier == "" {
		return nil, not_found.New("series", identifier)
	}

	return slices.Clone(c.Episodes[identifier]), nil
}
