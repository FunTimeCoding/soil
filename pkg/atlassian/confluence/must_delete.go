package confluence

import "github.com/funtimecoding/soil/pkg/errors"

func (c *Client) MustDelete(pageIdentifier string) {
	errors.PanicOnError(c.Delete(pageIdentifier))
}
