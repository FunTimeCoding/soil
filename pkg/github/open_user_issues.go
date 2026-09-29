package github

import "github.com/funtimecoding/soil/pkg/github/issue"

func (c *Client) OpenUserIssues() ([]*issue.Issue, error) {
	u, e := c.User()

	if e != nil {
		return nil, e
	}

	return c.SearchIssue("is:open is:issue author:%s", u.Name)
}
