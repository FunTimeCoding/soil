package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"github.com/funtimecoding/soil/pkg/web/stall_client/pool"
)

func wrap(
	l *locator.Locator,
	token string,
	p *pool.Pool,
) *Client {
	c, e := client.NewClientWithResponses(
		l.String(),
		client.WithHTTPClient(p.Client()),
		client.WithRequestEditorFn(web.BearerEditor(token)),
	)
	errors.PanicOnError(e)

	return &Client{client: c, base: l, token: token, pool: p}
}
