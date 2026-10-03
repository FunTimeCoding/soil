package jira

import (
	"github.com/funtimecoding/soil/pkg/atlassian/constant"
	"github.com/funtimecoding/soil/pkg/console"
)

func (c *Client) Addons() error {
	status, body, e := c.basic.GetPath(constant.JiraAddon)

	if e != nil {
		return e
	}

	console.Format("Addon: %d %s\n", status, body)

	return nil
}
