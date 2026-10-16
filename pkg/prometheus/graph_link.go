package prometheus

import "github.com/funtimecoding/soil/pkg/prometheus/constant"

func (c *Client) GraphLink(expression string) string {
	return c.link.Copy().Path(constant.Graph).Set(
		constant.Graph0Expression,
		expression,
	).String()
}
