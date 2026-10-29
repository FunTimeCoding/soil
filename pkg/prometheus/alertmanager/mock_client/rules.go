package mock_client

import "github.com/funtimecoding/soil/pkg/prometheus/rule/rule_list"

func (c *Client) Rules() (*rule_list.List, error) {
	return nil, c.failure
}
