package mock_client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/prometheus/rule/rule_list"
)

func (c *Client) MustSimpleSilence(
	alert string,
	r *rule_list.List,
) string {
	result, e := c.SimpleSilence(alert, r)
	errors.PanicOnError(e)

	return result
}
