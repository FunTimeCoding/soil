package client

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gojellyfind/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"github.com/funtimecoding/soil/pkg/web/stall_client"
	"net/http"
)

func newOnToken(
	l *locator.Locator,
	token func() string,
) *Client {
	p := stall_client.New().Traced().Build()
	c, e := client.NewClientWithResponses(
		l.String(),
		client.WithHTTPClient(p.Client()),
		client.WithRequestEditorFn(
			func(
				_ context.Context,
				q *http.Request,
			) error {
				web.Bearer(q, token())

				return nil
			},
		),
	)
	errors.PanicOnError(e)

	return &Client{client: c, base: l, token: token, pool: p}
}
