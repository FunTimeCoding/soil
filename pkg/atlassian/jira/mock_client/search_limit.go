package mock_client

import "github.com/funtimecoding/soil/pkg/atlassian/jira/issue"

func (c *Client) SearchLimit(
	int,
	string,
	...any,
) ([]*issue.Issue, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	if e := c.failures.Caller(); e != nil {
		return nil, e
	}

	return nil, nil
}
