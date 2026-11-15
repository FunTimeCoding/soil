package mock_client

import "github.com/funtimecoding/soil/pkg/atlassian/jira/link_type"

func (c *Client) LinkTypes() ([]link_type.Type, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	if e := c.failures.Caller(); e != nil {
		return nil, e
	}

	return nil, nil
}
