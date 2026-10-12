package openbao

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/openbao/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"github.com/openbao/openbao/api/v2"
)

func New(l *locator.Locator) *Client {
	c := api.DefaultConfig()
	c.Address = l.String()
	c.Timeout = constant.RequestTimeout
	client, e := api.NewClient(c)
	errors.PanicOnError(e)

	return &Client{client: client}
}
