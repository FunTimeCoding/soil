package github

import "github.com/funtimecoding/soil/pkg/github/issue"

func (c *Client) OpenUserIssues() []*issue.Issue {
	return c.MustSearchIssue("is:open is:issue author:%s", c.MustUser().Name)
}
