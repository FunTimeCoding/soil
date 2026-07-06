package opsgenie

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/opsgenie/opsgenie-go-sdk-v2/service"
)

func (c *Client) Services() []service.Service {
	r, e := c.userClient.Service.List(c.context, &service.ListRequest{})
	errors.PanicOnError(e)

	return r.Services
}
