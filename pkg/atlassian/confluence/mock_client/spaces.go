package mock_client

import "github.com/funtimecoding/soil/pkg/atlassian/confluence/space"

func (c *Client) Spaces() ([]*space.Space, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	if e := c.failures.Caller(); e != nil {
		return nil, e
	}

	return nil, nil
}
