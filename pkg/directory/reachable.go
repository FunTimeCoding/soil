package directory

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors"
)

func (c *Client) Reachable(_ context.Context) error {
	connection, e := c.connect()

	if e != nil {
		return e
	}

	errors.LogClose(connection)

	return nil
}
