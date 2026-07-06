package alertmanager

import "github.com/funtimecoding/soil/pkg/prometheus/rule/rule_list"

func (c *Client) Rules() (*rule_list.List, error) {
	if c.rules != nil {
		return c.rules, nil
	}

	v, e := c.prometheus.Rules()

	if e != nil {
		return nil, e
	}

	c.rules = v

	return c.rules, nil
}
