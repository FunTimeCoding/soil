package jira

import "github.com/funtimecoding/soil/pkg/errors"

func (c *Client) MustDynamicModules() {
	errors.PanicOnError(c.DynamicModules())
}
