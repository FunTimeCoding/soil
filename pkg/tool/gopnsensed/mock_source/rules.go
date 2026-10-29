package mock_source

import "github.com/funtimecoding/soil/pkg/opnsense/rule"

func (c *Client) Rules(_ string) ([]*rule.Rule, error) {
	return nil, nil
}
