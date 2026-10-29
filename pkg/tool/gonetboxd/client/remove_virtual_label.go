package client

import "context"

func (c *Client) RemoveVirtualLabel(
	x context.Context,
	machine string,
	key string,
) error {
	r, e := c.client.RemoveVirtualLabelWithResponse(x, machine, key)

	if e != nil {
		return e
	}

	return noContent("remove virtual label", r.HTTPResponse, r.Body)
}
