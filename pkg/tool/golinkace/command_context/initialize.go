package command_context

import (
	"github.com/funtimecoding/soil/pkg/tool/golinkaced/client"
	"github.com/funtimecoding/soil/pkg/web/locator"
)

func (c *Context) Initialize(
	host string,
	port int,
	insecure bool,
	token string,
) {
	l := locator.New(host)

	if port != 0 {
		l.Port(port)
	}

	if insecure {
		l.Insecure()
	}

	c.client = client.New(l, token)
}
