package jira

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/constant"
	"github.com/funtimecoding/soil/pkg/web/parameter"
)

func (c *Client) IssueV3(key string) error {
	status, body, e := c.basic.Get(
		c.basic.Base().Copy().Base(
			constant.Base,
		).Path("%s/%s", constant.Issue, key).Set(
			parameter.Fields,
			constant.AllFields,
		).String(),
	)

	if e != nil {
		return e
	}

	fmt.Printf("Response: %d %s", status, body)

	return nil
}
