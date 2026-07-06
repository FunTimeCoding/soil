package jira

import "github.com/funtimecoding/soil/pkg/errors"

func (c *Client) MustUnassign(key string) {
	errors.PanicOnError(c.Unassign(key))
}
