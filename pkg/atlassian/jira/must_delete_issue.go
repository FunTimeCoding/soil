package jira

import "github.com/funtimecoding/soil/pkg/errors"

func (c *Client) MustDeleteIssue(key string) {
	errors.PanicOnError(c.DeleteIssue(key))
}
