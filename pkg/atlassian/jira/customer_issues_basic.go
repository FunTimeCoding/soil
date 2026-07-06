package jira

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/constant"
	"github.com/funtimecoding/soil/pkg/web/parameter"
)

func (c *Client) CustomerIssuesBasic() error {
	status, body, e := c.basic.Get(
		c.basic.Base().Copy().Base(
			constant.ServiceDesk,
		).Path(constant.Request).SetInteger(
			parameter.Limit,
			10,
		).SetInteger(parameter.Start, 0).String(),
	)

	if e != nil {
		return e
	}

	fmt.Printf("Basic response: %d %s", status, body)

	return nil
}
