package graph

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors/not_configured"
	"github.com/funtimecoding/soil/pkg/microsoft/constant"
	"golang.org/x/oauth2"
	"time"
)

func (c *Client) current(renew bool) (*oauth2.Token, error) {
	c.lock.Lock()
	defer c.lock.Unlock()
	stored, found, e := c.store.FindToken()

	if e != nil {
		return nil, fmt.Errorf("load graph token: %w", e)
	}

	if !found {
		return nil, not_configured.Format(constant.SignInMissing)
	}

	if c.token == nil || stored.RefreshToken != c.stored {
		c.token = stored
		c.stored = stored.RefreshToken
	}

	if c.token.Valid() && !renew {
		return c.token, nil
	}

	stale := *c.token
	stale.Expiry = time.Unix(1, 0)
	t, f := c.configuration.TokenSource(c.tokenContext(), &stale).Token()

	if f != nil {
		return nil, refreshRefusal(f)
	}

	c.token = t

	if g := c.store.SaveToken(t); g != nil {
		return nil, fmt.Errorf("save graph token: %w", g)
	}

	c.stored = t.RefreshToken

	return t, nil
}
