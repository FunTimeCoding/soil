package jira

import (
	"github.com/funtimecoding/soil/pkg/atlassian/constant"
	"github.com/funtimecoding/soil/pkg/console"
	web "github.com/funtimecoding/soil/pkg/web/constant"
)

func (c *Client) CustomerIssuesBasic() error {
	status, body, e := c.basic.Get(
		c.basic.Base().Copy().Base(constant.JiraServiceDesk).Path(
			constant.JiraRequest,
		).SetInteger(
			web.ParameterLimit,
			10,
		).SetInteger(web.ParameterStart, 0).String(),
	)

	if e != nil {
		return e
	}

	console.Format("Basic response: %d %s", status, body)

	return nil
}
