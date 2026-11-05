package sublime

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
		c.base.Copy().Path("/views").String(),
		nil,
	)

	if e != nil {
		return e
	}

	r, f := c.client.Do(q)

	if f != nil {
		return fmt.Errorf("reach plugin: %w", f)
	}

	defer errors.LogClose(r.Body)

	if r.StatusCode != http.StatusOK {
		return statusError("reach plugin", r)
	}

	return nil
}
