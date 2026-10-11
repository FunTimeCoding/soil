package systemd

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/system/constant"
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
