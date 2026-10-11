package systemd

import "github.com/funtimecoding/soil/pkg/system/linux/systemd/status"

func (c *Client) Status(name string) *status.Result {
	return status.Parse(status.FilterLog(c.StatusOutput(name)))
}
