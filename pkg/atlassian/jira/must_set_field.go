package jira

import (
	"github.com/andygrunwald/go-jira"
	"github.com/funtimecoding/soil/pkg/errors"
)

func (c *Client) MustSetField(
	i *jira.Issue,
	name string,
	value any,
) {
	errors.PanicOnError(c.SetField(i, name, value))
}
