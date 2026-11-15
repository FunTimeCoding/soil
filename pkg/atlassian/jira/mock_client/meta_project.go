package mock_client

import (
	"github.com/andygrunwald/go-jira"
	"github.com/funtimecoding/soil/pkg/errors/not_found"
)

func (c *Client) MetaProject(key string) (*jira.MetaProject, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	if e := c.failures.Caller(); e != nil {
		return nil, e
	}

	return nil, not_found.New("project", key)
}
