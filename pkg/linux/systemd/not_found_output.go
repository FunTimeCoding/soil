package systemd

import (
	"github.com/funtimecoding/soil/pkg/linux/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
)

func (c *Client) NotFoundOutput() string {
	return c.ssh.Run(
		join.Space(
			constant.SystemdCommand,
			constant.SystemdState,
			constant.SystemdNotFound,
			constant.SystemdNoLegend,
		),
	).OutputString
}
