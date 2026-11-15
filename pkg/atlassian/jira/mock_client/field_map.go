package mock_client

import "github.com/funtimecoding/soil/pkg/atlassian/jira/field_map"

func (c *Client) FieldMap() (*field_map.Map, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	if e := c.failures.Caller(); e != nil {
		return nil, e
	}

	return nil, nil
}
