package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goalpined/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
)

func (c *Client) Packages(name string) string {
	parameters := &client.GetPackagesParams{}

	if name != "" {
		parameters.Name = &name
	}

	result, e := c.client.GetPackages(c.context, parameters)
	errors.PanicOnError(e)

	return web.ReadString(result)
}
