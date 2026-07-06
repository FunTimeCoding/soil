package jira

import "github.com/funtimecoding/soil/pkg/errors"

func (c *Client) MustIssueV3(key string) {
	errors.PanicOnError(c.IssueV3(key))
}
