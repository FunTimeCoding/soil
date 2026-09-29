package linkace

import "github.com/funtimecoding/soil/pkg/errors"

func (c *Client) MustDeleteNote(identifier int) {
	errors.PanicOnError(c.DeleteNote(identifier))
}
