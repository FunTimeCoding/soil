package mock_client

import "github.com/funtimecoding/soil/pkg/errors/sentry/basic/response"

func (c *Client) SearchIssues(
	_ string,
	_ string,
	_ string,
	_ int,
	_ string,
) ([]response.Issue, error) {
	return nil, nil
}
