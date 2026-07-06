package ollama

import "github.com/funtimecoding/soil/pkg/errors"

func (c *Client) MustHeartbeat() {
	errors.PanicOnError(c.Heartbeat())
}
