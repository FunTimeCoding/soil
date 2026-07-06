package alertmanager

import "github.com/funtimecoding/soil/pkg/errors"

func (c *Client) MustDeleteSilence(identifier string) {
	errors.PanicOnError(c.DeleteSilence(identifier))
}
