package openbao

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/openbao/constant"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"github.com/funtimecoding/soil/pkg/web/stall_client"
	"github.com/openbao/openbao/api/v2"
)

func New(l *locator.Locator) *Client {
	p := stall_client.New().Timeout(constant.RequestTimeout).Build()
	c := api.DefaultConfig()
	c.Address = l.String()
	c.Timeout = constant.RequestTimeout
	c.HttpClient = p.Halting()
	client, e := api.NewClient(c)
	errors.PanicOnError(e)

	return &Client{client: client, pool: p}
}
