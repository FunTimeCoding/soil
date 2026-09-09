package mock_client

import "github.com/funtimecoding/soil/pkg/errors/sentry/basic/response"

func (c *Client) IssueByIdentifier(
	_ string,
	_ string,
) (*response.Issue, error) {
	return nil, nil
}
