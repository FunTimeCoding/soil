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

	// 403 {"message":"Client must be authenticated as a system administrator to access this resource.","status-code":403}
	console.Format("Addon: %d %s\n", status, body)

	return nil
}
