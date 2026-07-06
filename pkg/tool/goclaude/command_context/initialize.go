package command_context

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
	"github.com/funtimecoding/soil/pkg/web/locator"
)

func (c *Context) Initialize(
	host string,
	port int,
) {
	r, e := client.NewClientWithResponses(
		locator.New(host).Port(port).String(),
	)
	errors.PanicOnError(e)
	c.client = r
}
