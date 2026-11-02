package alertmanager

import "github.com/funtimecoding/soil/pkg/prometheus/rule/rule_list"

func (c *Client) SimpleSilence(
	alert string,
	r *rule_list.List,
) (string, error) {
	return c.SetSilence(alert, "", 0, r)
}
