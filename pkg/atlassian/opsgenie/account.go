package opsgenie

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/opsgenie/opsgenie-go-sdk-v2/account"
)

func (c *Client) Account() *account.GetResult {
	result, e := c.userClient.Account.Get(c.context, &account.GetRequest{})
	errors.PanicOnError(e)

	return result
}
