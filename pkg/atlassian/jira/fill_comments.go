package jira

import (
	"github.com/funtimecoding/soil/pkg/atlassian/constant"
	"github.com/funtimecoding/soil/pkg/atlassian/jira/issue"
	"github.com/funtimecoding/soil/pkg/console"
)

func (c *Client) FillComments(v []*issue.Issue) error {
	filled := 0

	for _, i := range v {
		if i.Raw.Fields.Comments == nil {
			continue
		}

		if len(i.Raw.Fields.Comments.Comments) < constant.JiraCommentCap {
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
		console.Format("Filled comments for %d issue(s)\n", filled)
	}

	return nil
}
