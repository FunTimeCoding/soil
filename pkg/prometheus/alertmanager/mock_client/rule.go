package mock_client

import "github.com/funtimecoding/soil/pkg/prometheus/rule"

func (c *Client) Rule(_ string) (*rule.Rule, error) {
	return nil, c.failure
}
