package jira

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/atlassian/constant"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/basic/response"
)

func (c *Client) commentPage(
	key string,
	startAt int,
) (*response.Comments, error) {
	var result response.Comments

	if e := c.basic.Get(
		c.basic.Base().Copy().Base("/rest/api/2").Path(
			fmt.Sprintf("%s/%s/comment", constant.JiraIssue, key),
		).SetInteger(
			constant.JiraMaximumResultsKey,
			constant.JiraCommentPageSize,
		).SetInteger(
			"startAt",
			startAt,
		).String(),
		&result,
	); e != nil {
		return nil, e
	}

	return &result, nil
}
