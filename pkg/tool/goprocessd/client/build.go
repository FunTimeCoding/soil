package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goprocessd/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"github.com/funtimecoding/soil/pkg/web/stall_client/pool"
)

func build(
	l *locator.Locator,
	token string,
	p *pool.Pool,
) *Client {
	options := []client.ClientOption{client.WithHTTPClient(p.Client())}

	if token != "" {
		options = append(
			options,
			client.WithRequestEditorFn(web.BearerEditor(token)),
		)
	}

	c, e := client.NewClientWithResponses(l.String(), options...)
	errors.PanicOnError(e)

	return &Client{client: c, base: l, token: token, pool: p}
}
