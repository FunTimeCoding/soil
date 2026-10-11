package systemd

import "github.com/funtimecoding/soil/pkg/system/linux/systemd/show"

func (c *Client) Show(name string) *show.Result {
	return show.Parse(show.OutputToMap(c.ShowOutput(name)))
}
