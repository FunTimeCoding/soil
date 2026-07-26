package systemd

import (
	"github.com/funtimecoding/soil/pkg/linux/constant"
	"github.com/funtimecoding/soil/pkg/linux/systemd/show"
)

func (c *Client) IsActiveRunning(name string) bool {
	s := show.Parse(show.OutputToMap(c.ShowOutput(name)))

	if s.ActiveState == constant.SystemdActiveState &&
		s.SubState == constant.SystemdRunningSubState {
		return true
	}

	return false
}
