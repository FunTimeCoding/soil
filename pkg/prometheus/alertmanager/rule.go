package alertmanager

import "github.com/funtimecoding/soil/pkg/prometheus/rule"

func (c *Client) Rule(name string) (*rule.Rule, error) {
	v, e := c.Rules()

	if e != nil {
		return nil, e
	}

	return v.Find(name), nil
}
