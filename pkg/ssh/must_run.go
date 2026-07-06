package ssh

import "github.com/funtimecoding/soil/pkg/errors"

func (c *Client) MustRun(command string) string {
	r := c.Run(command)
	errors.PanicOnError(r.Error)

	return r.OutputString
}
