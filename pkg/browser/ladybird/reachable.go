package ladybird

import (
	"context"
	"github.com/funtimecoding/soil/pkg/browser/constant"
	"github.com/funtimecoding/soil/pkg/errors"
	"net/http"
)

func (c *Client) Reachable(x context.Context) error {
	q, e := http.NewRequestWithContext(
		x,
		http.MethodGet,
		c.driverLocator(constant.StatusRoute),
		nil,
	)

	if e != nil {
		return e
	}

	q.Close = true
	r, f := c.pool.Client().Do(q)

	if f != nil {
		return f
	}

	errors.LogClose(r.Body)

	return nil
}
