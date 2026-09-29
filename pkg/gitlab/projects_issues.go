package gitlab

import "github.com/funtimecoding/soil/pkg/gitlab/issue"

func (c *Client) ProjectsIssues() ([]*issue.Issue, error) {
	var result []*issue.Issue

	for _, identifier := range c.projects {
		issues, e := c.ProjectIssues(identifier)

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
