package kestra

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/kestra/constant"
)

func (c *Client) Namespaces() []string {
	result, r, e := c.client.FlowsAPI.ListDistinctNamespaces(
		c.context,
		constant.MainTenant,
	).Execute()
	errors.PanicOnWebError(r, e)

	return result
}
