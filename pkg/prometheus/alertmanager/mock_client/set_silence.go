package mock_client

import (
	"github.com/funtimecoding/soil/pkg/prometheus/rule/rule_list"
	"time"
)

func (c *Client) SetSilence(
	_ string,
	_ string,
	_ time.Duration,
	_ *rule_list.List,
) (string, error) {
	return "", c.failure
}
