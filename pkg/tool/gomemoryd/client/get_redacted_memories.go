package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/web/refusal"
)

func (c *Client) GetRedactedMemories(
	x context.Context,
) (map[int64]bool, error) {
	r, e := c.client.GetRedactedMemoriesWithResponse(x)

	if e != nil {
		return nil, e
	}

	if r.JSON200 == nil {
		return nil, refusal.New("redacted memories", r.HTTPResponse, r.Body)
	}

	result := map[int64]bool{}

	for _, identifier := range *r.JSON200 {
		result[identifier] = true
	}

	return result, nil
}
