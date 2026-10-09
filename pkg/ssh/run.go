package ssh

import (
	"github.com/funtimecoding/soil/pkg/ssh/command"
	"github.com/funtimecoding/soil/pkg/system/result"
)

func (c *Client) Run(s string) *result.Result {
	return c.RunCommand(command.New(s))
}
