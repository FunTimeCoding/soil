package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goproxmoxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/funtimecoding/soil/pkg/web/locator"
)

func New(
	l *locator.Locator,
	instance string,
	token string,
) *Client {
	c, e := client.NewClientWithResponses(
		l.String(),
		client.WithRequestEditorFn(web.BearerEditor(token)),
	)
	errors.PanicOnError(e)

	return &Client{
		context:  context.Background(),
		client:   c,
		instance: instance,
	}
}
