package jira

import (
	"github.com/andygrunwald/go-jira"
	"github.com/funtimecoding/soil/pkg/errors"
)

func (c *Client) MustCreateMeta(key string) *jira.CreateMetaInfo {
	result, e := c.CreateMeta(key)
	errors.PanicOnError(e)

	return result
}
