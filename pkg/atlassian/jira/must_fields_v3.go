package jira

import "github.com/funtimecoding/soil/pkg/errors"

func (c *Client) MustFieldsV3() {
	errors.PanicOnError(c.FieldsV3())
}
