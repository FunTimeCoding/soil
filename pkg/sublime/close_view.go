package sublime

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/system"
	"net/http"
)

func (c *Client) CloseView(identifier int) error {
	q, e := http.NewRequest(
		http.MethodDelete,
		c.base.Copy().Path("/views/%d", identifier).String(),
		nil,
	)

	if e != nil {
		return fmt.Errorf("close view: %w", e)
	}

	r, f := c.client.Do(q)

	if f != nil {
		return fmt.Errorf("close view: %w", f)
	}

	defer errors.LogClose(r.Body)

	if r.StatusCode != http.StatusOK {
		return fmt.Errorf(
			"close view: %d: %s",
			r.StatusCode,
			system.ReadAll(r.Body),
		)
	}

	return nil
}
