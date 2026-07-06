package ssh

import (
	"github.com/funtimecoding/soil/pkg/ssh/result"
	"github.com/funtimecoding/soil/pkg/strings/trim"
	"github.com/funtimecoding/soil/pkg/system/secure_shell"
)

func (c *Client) Run(command string) *result.Result {
	s := secure_shell.Session(c.client)
	defer secure_shell.CloseSession(s)
	stdout, stderr := secure_shell.SessionBuffers(s)
	e := s.Run(command)
	r := result.New(
		trim.NewLine(stdout.String()),
		trim.NewLine(stderr.String()),
		secure_shell.Exit(e),
		e,
	)

	if c.Panic {
		result.PanicOnExit(r)
		result.PanicOnError(r)
	}

	return r
}
