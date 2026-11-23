package graph

import (
	"context"
	"fmt"
	"golang.org/x/oauth2"
)

func (c *Client) Login(
	x context.Context,
	show func(*oauth2.DeviceAuthResponse),
) error {
	x = context.WithValue(x, oauth2.HTTPClient, c.pool.Client())
	a, e := c.configuration.DeviceAuth(x)

	if e != nil {
		return fmt.Errorf("request device code: %w", e)
	}

	show(a)
	t, e := c.configuration.DeviceAccessToken(x, a)

	if e != nil {
		return loginRefusal(e)
	}

	c.lock.Lock()
	defer c.lock.Unlock()
	c.token = t

	if f := c.store.SaveToken(t); f != nil {
		return fmt.Errorf("save graph token: %w", f)
	}

	return nil
}
