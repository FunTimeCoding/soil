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

	if c.token == nil {
		t, found, e := c.store.FindToken()

		if e != nil {
			return nil, fmt.Errorf("load graph token: %w", e)
		}

		if !found {
			return nil, not_configured.Format(constant.SignInMissing)
		}

		c.token = t
	}

	if c.token.Valid() && !renew {
		return c.token, nil
	}

	stale := *c.token
	stale.Expiry = time.Unix(1, 0)
	t, e := c.configuration.TokenSource(c.tokenContext(), &stale).Token()

	if e != nil {
		return nil, refreshRefusal(e)
	}

	c.token = t

	if f := c.store.SaveToken(t); f != nil {
		return nil, fmt.Errorf("save graph token: %w", f)
	}

	return t, nil
}
