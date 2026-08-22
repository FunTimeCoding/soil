package jira

import (
	"github.com/andygrunwald/go-jira"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
)

func (c *Client) MetaProject(key string) (*jira.MetaProject, error) {
	meta, e := c.CreateMeta(key)

	if e != nil {
		return nil, e
	}

	result := meta.GetProjectWithKey(key)

	if result == nil {
		return nil, not_found.New("project", key)
	}

	return result, nil
}
