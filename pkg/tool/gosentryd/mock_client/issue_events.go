package mock_client

import "github.com/funtimecoding/soil/pkg/errors/sentry/basic/response"

func (c *Client) IssueEvents(
	_ string,
	_ string,
	_ string,
	_ int,
	_ string,
) ([]response.Event, error) {
	return nil, nil
}
