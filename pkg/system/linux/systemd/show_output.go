package systemd

import (
	"github.com/funtimecoding/soil/pkg/strings/join"
	"github.com/funtimecoding/soil/pkg/system/constant"
)

func (c *Client) ShowOutput(name string) string {
	return c.ssh.Run(
		join.Space(constant.SystemdCommand, constant.SystemdShow, name),
	).OutputString
}
