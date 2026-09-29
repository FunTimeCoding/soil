package linkace

import "github.com/funtimecoding/soil/pkg/errors"

func (c *Client) MustDeleteTag(identifier int) {
	errors.PanicOnError(c.DeleteTag(identifier))
}
