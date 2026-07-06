package mock_jira

import "github.com/funtimecoding/soil/pkg/atlassian/jira/issue"

func (c *Client) SearchLimit(
	int,
	string,
	...any,
) ([]*issue.Issue, error) {
	return nil, nil
}
