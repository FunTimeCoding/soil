package mock_client

import "github.com/funtimecoding/soil/pkg/atlassian/jira/link_type"

func (c *Client) LinkTypes() ([]link_type.Type, error) {
	return nil, nil
}
