package jira

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/constant"
)

func (c *Client) DynamicModules() error {
	status, body, e := c.basic.GetPath(constant.Dynamic)

	if e != nil {
		return e
	}

	// 401 {"message":"The request is not from a Connect app."}
	fmt.Printf("DynamicModule: %d %s\n", status, body)

	return nil
}
