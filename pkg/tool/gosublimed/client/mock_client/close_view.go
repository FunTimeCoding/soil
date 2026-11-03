package mock_client

import (
	"context"
	"slices"
)

func (c *Client) CloseView(
	_ context.Context,
	identifier int,
) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if c.Failure != nil {
		return c.Failure
	}

	if e := c.failures.Caller(); e != nil {
		return e
	}

	i, e := c.index(identifier)

	if e != nil {
		return e
	}

	c.ViewList = slices.Delete(c.ViewList, i, i+1)

	return nil
}
