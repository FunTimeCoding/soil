package systemd

import (
	"github.com/funtimecoding/soil/pkg/linux/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
)

func (c *Client) StatusOutput(name string) string {
	return c.ssh.Run(
		join.Space(
			constant.SystemdCommand,
			constant.SystemdStatus,
			constant.SystemdOutput,
			constant.SystemdNotation,
			name,
		),
	).OutputString
}
