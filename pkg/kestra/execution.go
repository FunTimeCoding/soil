package kestra

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/kestra/constant"
	"github.com/kestra-io/client-sdk/go-sdk/kestra_api_client"
)

func (c *Client) Execution(execution string) *kestra_api_client.Execution {
	result, r, e := c.client.ExecutionsAPI.Execution(
		c.context,
		execution,
		constant.MainTenant,
	).Execute()
	errors.PanicOnWebError(r, e)

	return result
}
