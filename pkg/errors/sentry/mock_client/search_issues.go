package mock_client

import "github.com/funtimecoding/soil/pkg/errors/sentry/basic/response"

func (c *Client) SearchIssues(
	_ string,
	_ string,
	project string,
	_ int,
	_ string,
) ([]response.Issue, error) {
	c.issueSearchProject = project

	if c.Failure != nil {
		return nil, c.Failure
	}

	var result []response.Issue

	for _, i := range c.issues {
		if project == "" || i.Project.Identifier == project {
			result = append(result, *i)
		}
	}

	return result, nil
}
