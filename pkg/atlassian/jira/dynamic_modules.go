package jira

import (
	"github.com/funtimecoding/soil/pkg/atlassian/constant"
	"github.com/funtimecoding/soil/pkg/console"
)

func (c *Client) DynamicModules() error {
	status, body, e := c.basic.GetPath(constant.JiraDynamic)

	if e != nil {
		return e
	}

	console.Format("DynamicModule: %d %s\n", status, body)

	return nil
}
