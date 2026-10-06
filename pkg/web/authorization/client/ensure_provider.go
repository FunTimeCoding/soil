package client

import (
	"context"
	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) ensureProvider() error {
	c.providerMutex.Lock()
	defer c.providerMutex.Unlock()

	if c.provider != nil {
		return nil
	}

	p, e := oidc.NewProvider(
		oidc.ClientContext(context.Background(), web.StallClient()),
		c.issuer,
	)

	if e != nil {
		return e
	}

	c.provider = p
	c.verifier = p.Verifier(&oidc.Config{ClientID: c.identifier})

	return nil
}
