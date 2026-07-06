package jira

import (
	"github.com/andygrunwald/go-jira"
	"github.com/funtimecoding/soil/pkg/errors"
)

func (c *Client) MustEditMeta(key string) *jira.EditMetaInfo {
	result, e := c.EditMeta(key)
	errors.PanicOnError(e)

	return result
}
