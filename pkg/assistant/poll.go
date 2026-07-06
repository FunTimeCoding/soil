package assistant

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"time"
)

func (c *Client) poll() {
	t := time.NewTicker(30 * time.Second)
	defer t.Stop()

	for range t.C {
		errors.PanicOnError(c.connection.Ping())
	}
}
