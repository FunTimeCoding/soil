package ssh

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/ssh/command"
	"github.com/funtimecoding/soil/pkg/ssh/constant"
	"github.com/funtimecoding/soil/pkg/ssh/result"
	"github.com/funtimecoding/soil/pkg/strings/join/key_value"
	"github.com/funtimecoding/soil/pkg/strings/trim"
	"github.com/funtimecoding/soil/pkg/system/secure_shell"
	"golang.org/x/crypto/ssh"
)

func (c *Client) RunCommand(o *command.Command) *result.Result {
	s := secure_shell.Session(c.client)
	defer secure_shell.CloseSession(s)
	stdout, stderr := secure_shell.SessionBuffers(s)

	if o.RequestTeletype {
		errors.PanicOnError(
			s.RequestPty(
				"xterm",
				25,
				80,
				ssh.TerminalModes{
					ssh.TTY_OP_ISPEED: constant.TerminalBaudRate,
					ssh.TTY_OP_OSPEED: constant.TerminalBaudRate,
				},
			),
		)
	}

	var text string

	if prefix := EnvironmentPrefix(o); prefix != "" {
		text = key_value.Space(prefix, o.Command)
	} else {
		text = o.Command
	}

	e := s.Run(text)

	return result.New(
		trim.NewLine(stdout.String()),
		trim.NewLine(stderr.String()),
		secure_shell.Exit(e),
		e,
	)
}
