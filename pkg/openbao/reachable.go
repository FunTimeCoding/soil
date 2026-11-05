package openbao

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors/unavailable"
	"github.com/funtimecoding/soil/pkg/openbao/constant"
)

func (c *Client) Reachable(x context.Context) error {
	if e := c.authenticate(x); e != nil {
		return e
	}

	s, e := c.Health()

	if e != nil {
		return e
	}

	if s.Sealed || !s.Initialized {
		return unavailable.Format(constant.NotReady, c.Address())
	}

	return nil
}
