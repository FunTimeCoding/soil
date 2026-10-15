package inventory

import (
	"context"
	"github.com/funtimecoding/soil/pkg/errors"
	"github.com/funtimecoding/soil/pkg/tool/gonetboxd/generated/client"
	"github.com/funtimecoding/soil/pkg/web"
	"github.com/funtimecoding/soil/pkg/web/locator"
	"github.com/funtimecoding/soil/pkg/web/stall_client"
)

func New(
	l *locator.Locator,
	token string,
) *Inventory {
	p := stall_client.New().Build()
	c, e := client.NewClientWithResponses(
		l.String(),
		client.WithHTTPClient(p.Client()),
		client.WithRequestEditorFn(web.BearerEditor(token)),
	)
	errors.PanicOnError(e)

	return &Inventory{context: context.Background(), client: c, pool: p}
}
