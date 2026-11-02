package mock_client

import "github.com/funtimecoding/soil/pkg/prometheus/rule/rule_list"

func (c *Client) SimpleSilence(
	_ string,
	_ *rule_list.List,
) (string, error) {
	return "", c.failure
}
