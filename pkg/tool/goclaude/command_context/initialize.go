package command_context

import (
	"github.com/funtimecoding/soil/pkg/web"
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

	c.client = connect(l.String(), web.StallClient(), token)
	c.longClient = connect(l.String(), web.LongStallClient(), token)
}
