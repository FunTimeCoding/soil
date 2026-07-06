package jira

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/constant"
)

func (c *Client) Addons() error {
	status, body, e := c.basic.GetPath(constant.Addon)

	if e != nil {
		return e
	}

	// 403 {"message":"Client must be authenticated as a system administrator to access this resource.","status-code":403}
	fmt.Printf("Addon: %d %s\n", status, body)

	return nil
}
