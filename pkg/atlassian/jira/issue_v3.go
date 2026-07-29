package jira

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/atlassian/constant"
	web "github.com/funtimecoding/soil/pkg/web/constant"
)

func (c *Client) IssueV3(key string) error {
	status, body, e := c.basic.Get(
		c.basic.Base().Copy().Base(
			constant.JiraBase,
		).Path("%s/%s", constant.JiraIssue, key).Set(
			web.ParameterFields,
			constant.JiraAllFields,
		).String(),
	)

	if e != nil {
		return e
	}

	fmt.Printf("Response: %d %s", status, body)

	return nil
}
