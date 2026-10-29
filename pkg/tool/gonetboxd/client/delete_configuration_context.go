package client

import "context"

func (c *Client) DeleteConfigurationContext(
	x context.Context,
	name string,
) error {
	r, e := c.client.DeleteConfigurationContextWithResponse(x, name)

	if e != nil {
		return e
	}

	return noContent("delete configuration context", r.HTTPResponse, r.Body)
}
