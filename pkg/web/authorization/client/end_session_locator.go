package client

import "github.com/funtimecoding/soil/pkg/web/authorization/client/response"

func (c *Client) endSessionLocator() string {
	if c.ensureProvider() != nil {
		return ""
	}

	var m response.Discovery

	if e := c.provider.Claims(&m); e != nil {
		return ""
	}

	return m.EndSessionLocator
}
