package client

import (
	"context"
	"github.com/coreos/go-oidc/v3/oidc"
)

func (c *Client) ensureProvider() error {
	c.providerMutex.Lock()
	defer c.providerMutex.Unlock()

	if c.provider != nil {
		return nil
	}

	p, e := oidc.NewProvider(
		oidc.ClientContext(context.Background(), c.pool.Client()),
		c.issuer,
	)

	if e != nil {
		return e
	}

	c.provider = p
	c.verifier = p.Verifier(&oidc.Config{ClientID: c.identifier})

	return nil
}
