package openbao

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"github.com/openbao/openbao/api/v2"
	"time"
)

func New(host string) *Client {
	c := api.DefaultConfig()
	c.Address = locator.New(host).String()
	c.Timeout = 10 * time.Second
	client, e := api.NewClient(c)
	errors.PanicOnError(e)

	return &Client{client: client}
}
