package iterm

import (
	"context"
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"net/http"
)

func (c *Client) Reachable(x context.Context) error {
	q, e := http.NewRequestWithContext(
		x,
		http.MethodGet,
		c.base.Copy().Path("/sessions").String(),
		nil,
	)

	if e != nil {
		return e
	}

	r, f := c.client.Do(q)

	if f != nil {
		return fmt.Errorf("reach bridge: %w", f)
	}

	defer errors.LogClose(r.Body)

	if r.StatusCode != http.StatusOK {
		return statusError("reach bridge", r)
	}

	return nil
}
