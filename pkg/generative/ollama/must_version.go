package ollama

import "github.com/funtimecoding/soil/pkg/errors"

func (c *Client) MustVersion() string {
	result, e := c.Version()
	errors.PanicOnError(e)

	return result
}
