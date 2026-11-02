package gitlab

import (
	"context"
	"github.com/funtimecoding/soil/pkg/gitlab/issue"
)

func (c *Client) ProjectsIssues(x context.Context) ([]*issue.Issue, error) {
	var result []*issue.Issue

	for _, identifier := range c.projects {
		issues, e := c.ProjectIssues(x, identifier)

		if e != nil {
			return nil, e
		}

		for _, i := range issues {
			if i.Done() {
				continue
			}

			result = append(result, i)
		}
	}

	return result, nil
}
