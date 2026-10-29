package client

import "context"

func (c *Client) ReloadProcfile(x context.Context) error {
	r, e := c.client.ReloadProcfileWithResponse(x)

	if e != nil {
		return e
	}

	return noContent("reload procfile", r.HTTPResponse, r.Body)
}
