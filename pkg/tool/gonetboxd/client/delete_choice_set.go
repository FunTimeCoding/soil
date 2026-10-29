package client

import "context"

func (c *Client) DeleteChoiceSet(
	x context.Context,
	name string,
) error {
	r, e := c.client.DeleteChoiceSetWithResponse(x, name)

	if e != nil {
		return e
	}

	return noContent("delete choice set", r.HTTPResponse, r.Body)
}
