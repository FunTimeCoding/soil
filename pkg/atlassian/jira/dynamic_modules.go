package jira

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/atlassian/constant"
)

func (c *Client) DynamicModules() error {
	status, body, e := c.basic.GetPath(constant.JiraDynamic)

	if e != nil {
		return e
	}

	// 401 {"message":"The request is not from a Connect app."}
	fmt.Printf("DynamicModule: %d %s\n", status, body)

	return nil
}
