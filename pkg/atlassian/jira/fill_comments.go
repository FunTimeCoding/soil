package jira

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/constant"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/issue"
)

func (c *Client) FillComments(v []*issue.Issue) error {
	filled := 0

	for _, i := range v {
		if i.Raw.Fields.Comments == nil {
			continue
		}

		if len(i.Raw.Fields.Comments.Comments) < constant.CommentCap {
			continue
		}

		all, e := c.allComments(i.Key)

		if e != nil {
			return e
		}

		i.Raw.Fields.Comments.Comments = all
		filled++
	}

	if filled > 0 {
		fmt.Printf(
			"Filled comments for %d issue(s)\n",
			filled,
		)
	}

	return nil
}
