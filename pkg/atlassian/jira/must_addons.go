package jira

import "github.com/funtimecoding/soil/pkg/errors"

func (c *Client) MustAddons() {
	errors.PanicOnError(c.Addons())
}
