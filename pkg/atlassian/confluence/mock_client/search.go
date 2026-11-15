package mock_client

import "github.com/funtimecoding/soil/pkg/atlassian/confluence/search_result"

func (c *Client) Search(
	query string,
	a ...any,
) ([]*search_result.Result, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	if e := c.failures.Caller(); e != nil {
		return nil, e
	}

	return nil, nil
}
