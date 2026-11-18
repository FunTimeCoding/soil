package nut

import "context"

func (c *Client) Reachable(x context.Context) error {
	_, e := c.Batteries(x)

	return e
}
