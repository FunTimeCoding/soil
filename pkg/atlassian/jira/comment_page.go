package jira

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/basic/response"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/constant"
	"github.com/funtimecoding/soil/pkg/notation"
)

func (c *Client) commentPage(
	key string,
	startAt int,
) (*response.Comments, error) {
	var result response.Comments
	_, r, e := c.basic.Get(
		c.basic.Base().Copy().Base("/rest/api/2").Path(
			fmt.Sprintf("%s/%s/comment", constant.Issue, key),
		).SetInteger(
			constant.MaximumResultsKey,
			constant.CommentPageSize,
		).SetInteger(
			"startAt",
			startAt,
		).String(),
	)

	if e != nil {
		return nil, e
	}

	notation.MustDecode(r, &result, true)

	return &result, nil
}
