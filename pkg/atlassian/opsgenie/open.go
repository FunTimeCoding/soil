package opsgenie

import "github.com/funtimecoding/soil/pkg/atlassian/opsgenie/alert"

func (c *Client) Open() []*alert.Alert {
	return c.Search("status:%s", alert.OpenStatus)
}
