package client

import (
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/goqueryd/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"github.com/funtimecoding/soil/pkg/web/stall_client"
)

func New(
	l *locator.Locator,
	token string,
	o ...Option,
) *Client {
	result := &Client{base: l, token: token}

	for _, f := range o {
		f(result)
	}

	b := stall_client.New().Traced()

	if result.long {
		b.Long()
	}

	result.pool = b.Build()
	c, e := client.NewClientWithResponses(
		l.String(),
		client.WithHTTPClient(result.pool.Client()),
		client.WithRequestEditorFn(web.BearerEditor(token)),
	)
	errors.PanicOnError(e)
	result.client = c

	return result
}
