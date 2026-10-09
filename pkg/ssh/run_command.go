package ssh

import (
	"github.com/funtimecoding/soil/pkg/errors"
	processError "github.com/funtimecoding/soil/pkg/errors/command"
	"github.com/funtimecoding/soil/pkg/ssh/command"
	"github.com/funtimecoding/soil/pkg/ssh/constant"
	"github.com/funtimecoding/soil/pkg/strings/join/key_value"
	"github.com/funtimecoding/soil/pkg/strings/trim"
	"github.com/funtimecoding/soil/pkg/system/result"
	"github.com/funtimecoding/soil/pkg/system/secure_shell"
	"golang.org/x/crypto/ssh"
)

func (c *Client) RunCommand(o *command.Command) *result.Result {
	s := secure_shell.Session(c.dialed())
	defer secure_shell.CloseSession(s)
	stdout, stderr := secure_shell.SessionBuffers(s)

	if o.RequestTeletype {
		errors.PanicOnError(
			s.RequestPty(
				constant.TerminalType,
				constant.TerminalHeight,
				constant.TerminalWidth,
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
	r := result.New(
		trim.NewLine(stdout.String()),
		trim.NewLine(stderr.String()),
		secure_shell.Exit(e),
		nil,
	)

	if e != nil {
		r.Error = processError.New(text, r.OutputString, r.ErrorString, e)
	}

	if c.Panic {
		errors.PanicOnError(r.Error)
	}

	return r
}
