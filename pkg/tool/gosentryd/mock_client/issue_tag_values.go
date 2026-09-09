package mock_client

import "github.com/funtimecoding/soil/pkg/errors/sentry/basic/response"

func (c *Client) IssueTagValues(
	_ string,
	_ string,
	_ string,
	_ int,
) ([]response.TagValue, error) {
	return nil, nil
}
