package client

import "context"

func (c *Client) ReloadEnvironment(x context.Context) error {
	r, e := c.client.ReloadEnvironmentWithResponse(x)

	if e != nil {
		return e
	}

	return noContent("reload environment", r.HTTPResponse, r.Body)
}
