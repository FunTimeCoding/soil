package keepass

import (
	"errors"
	"github.com/funtimecoding/soil/pkg/errors/conflict"
)

func (c *Client) Save() error {
	if c.Changed() {
		return conflict.Format(
			"database changed on disk since load - refusing to save",
		)
	}

	if e := c.database.LockProtectedEntries(); e != nil {
		return e
	}

	return errors.Join(c.write(), c.database.UnlockProtectedEntries())
}
