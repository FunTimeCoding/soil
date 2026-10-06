package nextcloud

import "github.com/funtimecoding/soil/pkg/errors"

func (c *Client) MustStatus() {
	errors.PanicOnError(c.Status())
}
