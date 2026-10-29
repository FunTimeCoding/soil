package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gomattermostd/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"github.com/funtimecoding/soil/pkg/web/stall_client"
)

func New(
	l *locator.Locator,
	token string,
) *Client {
	p := stall_client.New().Build()
	c, e := client.NewClientWithResponses(
		l.String(),
		client.WithHTTPClient(p.Client()),
		client.WithRequestEditorFn(web.BearerEditor(token)),
	)
	errors.PanicOnError(e)

	return &Client{client: c, base: l, token: token, pool: p}
}
