package mock_client

import "github.com/funtimecoding/soil/pkg/atlassian/jira/issue"

func (c *Client) Search(
	string,
	...any,
) ([]*issue.Issue, error) {
	return c.Issues, nil
}
