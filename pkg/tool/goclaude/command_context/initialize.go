package command_context

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goclauded/generated/client"
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

	r, e := client.NewClientWithResponses(
		l.String(),
		client.WithRequestEditorFn(web.BearerEditor(token)),
	)
	errors.PanicOnError(e)
	c.client = r
}
