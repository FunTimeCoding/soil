package chromium

import (
	"github.com/funtimecoding/soil/pkg/errors/connection"
	"github.com/funtimecoding/soil/pkg/errors/unreachable"
)

func (c *Client) dialFailure(e error) error {
	if f := connection.Classify(e); f != nil {
		return f
	}

	return unreachable.Format("browser %s:%d: %v", c.host, c.port, e)
}
