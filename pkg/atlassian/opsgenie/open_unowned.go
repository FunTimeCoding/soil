package opsgenie

import "github.com/funtimecoding/soil/pkg/atlassian/opsgenie/alert"

func (c *Client) OpenUnowned() []*alert.Alert {
	return c.Search("status:%s owner:NULL", alert.OpenStatus)
}
