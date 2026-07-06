package kestra

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/kestra/constant"
	"github.com/kestra-io/client-sdk/go-sdk/kestra_api_client"
)

func (c *Client) Flows(namespace string) []kestra_api_client.Flow {
	result, r, e := c.client.FlowsAPI.ListFlowsByNamespace(
		c.context,
		namespace,
		constant.MainTenant,
	).Execute()
	errors.PanicOnWebError(r, e)

	return result
}
