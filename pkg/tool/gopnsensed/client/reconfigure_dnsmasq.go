package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/web/refusal"
	"net/http"
)

func (c *Client) ReconfigureDnsmasq(x context.Context) error {
	r, e := c.client.ReconfigureDnsmasqWithResponse(x)

	if e != nil {
		return e
	}

	if r.StatusCode() != http.StatusNoContent {
		return refusal.New("reconfigure dnsmasq", r.HTTPResponse, r.Body)
	}

	return nil
}
