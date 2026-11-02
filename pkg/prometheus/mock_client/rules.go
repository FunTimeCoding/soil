package mock_client

import "github.com/funtimecoding/soil/pkg/prometheus/rule/rule_list"

func (c *Client) Rules() (*rule_list.List, error) {
	if c.Failure != nil {
		return nil, c.Failure
	}

	return rule_list.New(), nil
}
