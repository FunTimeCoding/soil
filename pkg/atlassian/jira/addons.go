package jira

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/atlassian/constant"
)

func (c *Client) Addons() error {
	status, body, e := c.basic.GetPath(constant.JiraAddon)

	if e != nil {
		return e
	}

	// 403 {"message":"Client must be authenticated as a system administrator to access this resource.","status-code":403}
	fmt.Printf("Addon: %d %s\n", status, body)

	return nil
}
