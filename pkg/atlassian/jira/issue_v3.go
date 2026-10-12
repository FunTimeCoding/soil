package jira

import (
	"github.com/funtimecoding/soil/pkg/atlassian/constant"
	"github.com/funtimecoding/soil/pkg/console"
	webConstant "github.com/funtimecoding/soil/pkg/web/constant"
)

func (c *Client) IssueV3(key string) error {
	body, e := c.basic.Bytes(
		c.basic.Base().Copy().Base(constant.JiraBase).Path(
			"%s/%s",
			constant.JiraIssue,
			key,
		).Set(
			webConstant.ParameterFields,
			constant.JiraAllFields,
		).String(),
	)

	if e != nil {
		return e
	}

	console.Format("Response: %s", body)

	return nil
}
