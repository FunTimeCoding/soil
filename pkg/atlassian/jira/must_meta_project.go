package jira

import (
	"github.com/andygrunwald/go-jira"
	"github.com/funtimecoding/soil/pkg/errors"
)

func (c *Client) MustMetaProject(key string) *jira.MetaProject {
	result, e := c.MetaProject(key)
	errors.PanicOnError(e)

	return result
}
