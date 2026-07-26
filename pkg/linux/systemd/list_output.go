package systemd

import (
	"github.com/funtimecoding/soil/pkg/linux/constant"
	"github.com/funtimecoding/soil/pkg/strings/join"
)

func (c *Client) ListOutput() string {
	return c.ssh.Run(
		join.Space(
			constant.SystemdCommand,
			constant.SystemdListUnits,
			constant.SystemdType,
			constant.SystemdService,
			constant.SystemdAll,
			constant.SystemdFull,
			constant.SystemdPlain,
			constant.SystemdNoLegend,
		),
	).OutputString
}
