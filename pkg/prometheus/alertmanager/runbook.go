package alertmanager

import "github.com/funtimecoding/soil/pkg/prometheus/alertmanager/alert/rule_parser"

func (c *Client) Runbook(name string) (string, error) {
	v, e := c.Rules()

	if e != nil {
		return "", e
	}

	return rule_parser.New(v).Runbook(name), nil
}
