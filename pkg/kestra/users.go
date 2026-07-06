package kestra

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/kestra-io/client-sdk/go-sdk/kestra_api_client"
)

func (c *Client) Users() []kestra_api_client.IAMUserControllerApiUserSummary {
	result, r, e := c.client.UsersAPI.ListUsers(
		c.context,
	).Page(0).Size(10).Execute()
	errors.PanicOnWebError(r, e)

	return result.Results
}
